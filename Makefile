install:
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

postgres:
	docker run --name postgres -e POSTGRES_PASSWORD=password -p5432:5432 -d postgres

run:
	go run ./cmd/*.go
	
path=./db/migration
dsn=postgresql://postgres:password@127.0.0.1:5432/auction?sslmode=disable
name?=migrate

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres auction

dropdb:
	docker exec -it postgres dropdb --username=postgres auction

# usage: make migrate-create name=name
migrate-create:
	migrate create -ext sql -dir $(path) -seq $(name) -verbose

migrateup:
	migrate -path $(path) -database "$(dsn)" -verbose up

migrateup1:
	migrate -path $(path) -database "$(dsn)" -verbose up 1

migratedown:
	migrate -path $(path) -database "$(dsn)" -verbose down

migratedown1:
	migrate -path $(path) -database "$(dsn)" -verbose down 1

sqlc:
	sqlc generate

.PHONY: install postgres createdb dropdb run migrateup migrateup1 migratedown migratedown1 sqlc
