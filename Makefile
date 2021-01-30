dependencies:
	go mod vendor

start:
	docker-compose up -d

stop:
	docker-compose down

mongo:
	docker-compose up -d mongo

run:
	docker-compose up -d mongo
	go run cmd/server/main.go
