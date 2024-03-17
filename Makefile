postgres:
	docker run --name SimpleBank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=010220 -d postgres:16-alpine

createdb:
	docker exec -it SimpleBank createdb --username=root --owner=root SimpleBank

dropdb:
	docker exec -it SimpleBank dropdb SimpleBank

migrateup:
	migrate -path db/migration -database "postgresql://root:010220@localhost:5432/SimpleBank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:010220@localhost:5432/SimpleBank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
