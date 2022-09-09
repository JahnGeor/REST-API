.SILENT:

build:
	docker-compose build todo-app

run:
	docker-compose up todo-app

migrate:
	migrate -path ./schema -database 'postgres://postgres:admin@localhost:5436/postgres?sslmode=disable' up