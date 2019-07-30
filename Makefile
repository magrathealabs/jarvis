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

stop_travis_services:
	# Disable services enabled by default
	# http://docs.travis-ci.com/user/database-setup/#MySQL
	sudo /etc/init.d/mysql stop
	sudo /etc/init.d/postgresql stop