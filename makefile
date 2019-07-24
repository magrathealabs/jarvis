up: build
	docker-compose up

build:
	docker-compose build

dev: build
	docker-compose up -d

test: dev
	go test ./...