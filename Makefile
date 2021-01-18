dependencies:
	go mod vendor

start:
	docker-compose up -d

stop:
	docker-compose down

mongo:
	docker-compose up -d mongo
