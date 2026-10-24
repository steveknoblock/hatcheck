cd ~/projects/hatcheck/
git fetch
git checkout develop-go
git pull
go build ./...
set -a && source .env && set +a
go run ./server
