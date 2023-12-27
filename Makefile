postgres:
	sudo docker run --name postgres-new -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres
createdb:
	sudo docker exec -it postgres-new createdb --username=postgres --owner=postgres simplebank

dropdb:
	sudo docker exec -it postgres-new dropdb simplebank
migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simplebank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simplebank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test