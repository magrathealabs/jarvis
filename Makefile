up: stop build
	docker-compose up

build:
	docker-compose build

dev: build
	docker-compose up -d postgres

stop:
	docker-compose down

test: dev
	go test ./...