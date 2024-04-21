postgres:
	docker run -itd -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123456 -p 5432:5432 --mount source=simplebank,target=/data --name postgresql --network bank-network postgres

createdb:
	docker exec -it postgresql createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgresql dropdb -U postgres -W simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:123456@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:123456@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://postgres:123456@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:123456@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/EricUCL/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock