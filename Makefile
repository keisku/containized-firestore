up:
	docker-compose up -d --remove-orphans

down:
	docker-compose down

build:
	docker-compose build --no-cache

logs:
	docker-compose logs --tail=10 -f example

restart:
	docker-compose restart example

test:
	docker-compose exec example go test -cover -v `go list ./... | grep -v mock | grep -v proto`

dump:
	docker-compose exec example go run scripts/dump/main.go
