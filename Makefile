postgres:
	docker run --name simple_banq -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it simple_banq createdb --username=root --owner=root simple_banq

dropdb:
	docker exec -it simple_banq dropdb --username=root --owner=root simple_banq

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_banq?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_banq?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

.PHONY: createdb dropdb postgres migratedown migrateup sqlc test server

