up:
	docker-compose up -d --remove-orphans

down:
	docker-compose down

build:
	docker-compose build --no-cache

logs:
	docker-compose logs --tail=10 -f app

restart:
	docker-compose restart app

test:
	docker-compose exec app go test -cover -v `go list ./... | grep -v mock | grep -v proto`
