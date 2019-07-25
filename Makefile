up: stop build
	docker-compose up

build:
	docker-compose build

dev: build
	docker-compose up -d postgres rabbitmq

stop:
	docker-compose down

test: dev migrate
	go test ./...

migrate:
	docker-compose exec postgres psql -U postgres -c "create database test"
	docker-compose exec postgres psql -U postgres -c "create database development"
	docker-compose exec postgres psql -U postgres -c "create database production"
