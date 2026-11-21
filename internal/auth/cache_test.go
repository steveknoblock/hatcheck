package auth

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// fakeClock is a controllable time source for cache expiry tests.
type fakeClock struct {
	mu sync.Mutex
	t  time.Time
}

func (f *fakeClock) now() time.Time {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.t
}

func (f *fakeClock) advance(d time.Duration) {
	f.mu.Lock()
	f.t = f.t.Add(d)
	f.mu.Unlock()
}

func newTestCache(ttl time.Duration, max int) (*identityCache, *fakeClock) {
	clk := &fakeClock{t: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)}
	c := newIdentityCache(ttl, max)
	c.now = clk.now
	return c, clk
}

func TestCacheReusesSuccessfulValidation(t *testing.T) {
	c, _ := newTestCache(30*time.Second, 10)
	var calls int
	validate := func() (Identity, error) { calls++; return Identity{UserID: "u1"}, nil }

	for i := 0; i < 5; i++ {
		id, err := c.do("tok", validate)
		if err != nil || id.UserID != "u1" {
			t.Fatalf("call %d: got %+v, %v", i, id, err)
		}
	}
	if calls != 1 {
		t.Errorf("validate ran %d times, want 1", calls)
	}
}

func TestCacheExpiresAfterTTL(t *testing.T) {
	c, clk := newTestCache(30*time.Second, 10)
	var calls int
	validate := func() (Identity, error) { calls++; return Identity{UserID: "u1"}, nil }

	c.do("tok", validate)
	clk.advance(29 * time.Second)
	c.do("tok", validate)
	if calls != 1 {
		t.Fatalf("within TTL: validate ran %d times, want 1", calls)
	}
	clk.advance(2 * time.Second)
	c.do("tok", validate)
	if calls != 2 {
		t.Errorf("after TTL: validate ran %d times, want 2", calls)
	}
}

func TestCacheDoesNotCacheFailures(t *testing.T) {
	c, _ := newTestCache(30*time.Second, 10)
	var calls int
	validate := func() (Identity, error) { calls++; return Identity{}, errors.New("expired") }

	for i := 0; i < 3; i++ {
		if _, err := c.do("bad", validate); err == nil {
			t.Fatal("expected an error")
		}
	}
	if calls != 3 {
		t.Errorf("failed validations ran %d times, want 3 (never cached)", calls)
	}
}

func TestCacheKeepsTokensSeparate(t *testing.T) {
	c, _ := newTestCache(30*time.Second, 10)
	a, _ := c.do("tokA", func() (Identity, error) { return Identity{UserID: "alice"}, nil })
	b, _ := c.do("tokB", func() (Identity, error) { return Identity{UserID: "bob"}, nil })
	if a.UserID != "alice" || b.UserID != "bob" {
		t.Errorf("got %q and %q", a.UserID, b.UserID)
	}
	again, _ := c.do("tokA", func() (Identity, error) { return Identity{UserID: "WRONG"}, nil })
	if again.UserID != "alice" {
		t.Errorf("tokA after tokB: got %q, want alice", again.UserID)
	}
}

func TestConcurrentCallersShareOneValidation(t *testing.T) {
	c, _ := newTestCache(30*time.Second, 10)
	var calls atomic.Int32
	release := make(chan struct{})
	validate := func() (Identity, error) {
		calls.Add(1)
		<-release // hold the first call open so the others pile up behind it
		return Identity{UserID: "u1"}, nil
	}

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if id, err := c.do("tok", validate); err != nil || id.UserID != "u1" {
				t.Errorf("got %+v, %v", id, err)
			}
		}()
	}
	time.Sleep(50 * time.Millisecond)
	close(release)
	wg.Wait()
	if n := calls.Load(); n != 1 {
		t.Errorf("validate ran %d times for 8 concurrent callers, want 1", n)
	}
}

func TestCacheStaysBounded(t *testing.T) {
	c, _ := newTestCache(time.Hour, 5)
	for i := 0; i < 50; i++ {
		tok := string(rune('a'+i%26)) + string(rune('A'+i/26))
		c.do(tok, func() (Identity, error) { return Identity{UserID: tok}, nil })
	}
	if n := len(c.entries); n > 5 {
		t.Errorf("cache holds %d entries, max is 5", n)
	}
}

func TestCacheDropsExpiredEntriesBeforeEvictingFreshOnes(t *testing.T) {
	c, clk := newTestCache(10*time.Second, 2)
	c.do("old1", func() (Identity, error) { return Identity{UserID: "o1"}, nil })
	c.do("old2", func() (Identity, error) { return Identity{UserID: "o2"}, nil })
	clk.advance(11 * time.Second)
	c.do("fresh", func() (Identity, error) { return Identity{UserID: "f"}, nil })
	if _, ok := c.entries[cacheKey("fresh")]; !ok {
		t.Error("fresh entry missing")
	}
	if n := len(c.entries); n != 1 {
		t.Errorf("expected expired entries purged, have %d entries", n)
	}
}

func TestValidateSessionJWTUsesCache(t *testing.T) {
	var calls int
	c := &Client{
		identities: newIdentityCache(identityCacheTTL, identityCacheMax),
		validate: func(ctx context.Context, jwt string) (Identity, error) {
			calls++
			return Identity{UserID: "user-" + jwt}, nil
		},
	}
	for i := 0; i < 4; i++ {
		id, err := c.ValidateSessionJWT(context.Background(), "abc")
		if err != nil || id.UserID != "user-abc" {
			t.Fatalf("got %+v, %v", id, err)
		}
	}
	if calls != 1 {
		t.Errorf("Stytch validation ran %d times, want 1", calls)
	}
}

func TestValidateSessionJWTSurvivesCancelledFirstCaller(t *testing.T) {
	c := &Client{
		identities: newIdentityCache(identityCacheTTL, identityCacheMax),
		validate: func(ctx context.Context, jwt string) (Identity, error) {
			if ctx.Err() != nil {
				return Identity{}, ctx.Err()
			}
			return Identity{UserID: "u"}, nil
		},
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := c.ValidateSessionJWT(ctx, "tok"); err != nil {
		t.Errorf("a cancelled request context should not fail the shared validation: %v", err)
	}
}
