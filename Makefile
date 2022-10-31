copy:
	copy configs
migrate-up:
	migrate -path ./schema -database 'postgres://jahngeor:admin@localhost:5436/avito-finance?sslmode=disable' up
migrate-down:
	migrate -path ./schema -database 'postgres://jahngeor:admin@localhost:5436/avito-finance?sslmode=disable' down
migrate-create:
	migrate create -ext sql -dir ./schema -seq avito-finance
build:
	docker-compose build
run:
	docker-compose up
restart: copy build
	docker-compose up