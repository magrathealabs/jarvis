up: stop build
	docker-compose up

build:
	docker-compose build

dev: build
	docker-compose up -d postgres rabbitmq graphite

stop:
	docker-compose down

test: stop dev migrate
	go test ./...

migrate:
	sleep 5
	docker-compose exec postgres psql -U postgres -c "create database test"
	docker-compose exec postgres psql -U postgres -c "create database development"
	docker-compose exec postgres psql -U postgres -c "create database production"
