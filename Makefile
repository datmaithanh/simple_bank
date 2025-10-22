postgres:
	docker run --name postgres17 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17.5-alpine3.21

createdb: 
	docker exec -it postgres17 createdb --username=root --owner=root simple-bank

dropdb:
	docker exec -it postgres17 dropdb simple-bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple-bank?sslmode=disable" -verbose up
	
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple-bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

tosql:
	docker exec -it postgres17 psql -U root

test:
	go test -v -cover -count=1 ./...

.PHONY: createdb dropdb postgres migrateup migratedown sqlc tosql test