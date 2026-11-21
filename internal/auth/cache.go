package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"sync"
	"time"
)

// Why this cache exists
//
// Stytch session JWTs have a 5-minute lifetime. Once a token is older than
// that, local validation fails and the Stytch SDK falls back to a remote
// API call — on every request. Measured against a real Hatcheck, that made
// each request take ~120-160 ms instead of ~15 ms, and a UI page that fires
// a hundred requests (two per post) spend seconds on authentication alone.
//
// identityCache remembers a successful validation for a short time, so a
// burst of requests carrying the same token costs one validation instead of
// one per request.
//
// Trade-off: a session that Stytch revokes can still be accepted for up to
// identityCacheTTL afterwards. That is far shorter than the 5-minute window
// a locally validated JWT already has, and Hatcheck has no server-side
// logout (the client just drops its token), so nothing else is weakened.
// Failures are never cached, so a bad or expired token is re-checked every
// time.
const (
	identityCacheTTL = 30 * time.Second
	identityCacheMax = 1024
)

type cachedIdentity struct {
	identity Identity
	expires  time.Time
}

// inflightCall lets concurrent requests with the same uncached token share
// one validation instead of each calling Stytch.
type inflightCall struct {
	wg       sync.WaitGroup
	identity Identity
	err      error
}

type identityCache struct {
	mu       sync.Mutex
	entries  map[string]cachedIdentity
	inflight map[string]*inflightCall
	ttl      time.Duration
	max      int
	now      func() time.Time // replaceable in tests
}

func newIdentityCache(ttl time.Duration, max int) *identityCache {
	return &identityCache{
		entries:  make(map[string]cachedIdentity),
		inflight: make(map[string]*inflightCall),
		ttl:      ttl,
		max:      max,
		now:      time.Now,
	}
}

// cacheKey hashes the token so raw JWTs are not kept in memory as map keys.
func cacheKey(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

// do returns the cached identity for token if it is still fresh; otherwise
// it runs validate once (shared with any concurrent callers for the same
// token) and caches a successful result.
func (c *identityCache) do(token string, validate func() (Identity, error)) (Identity, error) {
	key := cacheKey(token)

	c.mu.Lock()
	if e, ok := c.entries[key]; ok {
		if c.now().Before(e.expires) {
			c.mu.Unlock()
			return e.identity, nil
		}
		delete(c.entries, key)
	}
	if call, ok := c.inflight[key]; ok {
		c.mu.Unlock()
		call.wg.Wait()
		return call.identity, call.err
	}
	call := &inflightCall{}
	call.wg.Add(1)
	c.inflight[key] = call
	c.mu.Unlock()

	// Always release waiters and clear the in-flight entry, even if
	// validate panics.
	defer func() {
		c.mu.Lock()
		delete(c.inflight, key)
		if call.err == nil && call.identity.UserID != "" {
			c.store(key, call.identity)
		}
		c.mu.Unlock()
		call.wg.Done()
	}()

	call.identity, call.err = validate()
	return call.identity, call.err
}

// store adds an entry, keeping the cache bounded. Callers hold c.mu.
func (c *identityCache) store(key string, id Identity) {
	if len(c.entries) >= c.max {
		now := c.now()
		for k, e := range c.entries {
			if !now.Before(e.expires) {
				delete(c.entries, k)
			}
		}
		for k := range c.entries { // still full: drop an arbitrary entry
			if len(c.entries) < c.max {
				break
			}
			delete(c.entries, k)
		}
	}
	c.entries[key] = cachedIdentity{identity: id, expires: c.now().Add(c.ttl)}
}
