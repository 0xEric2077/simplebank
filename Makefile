postgres:
	docker run -itd -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123456 -p 5432:5432 --mount source=simplebank,target=/data --name postgresql postgres

createdb:
	docker exec -it postgresql createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgresql dropdb -U postgres -W simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:123456@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:123456@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown