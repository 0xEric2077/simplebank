package main

import (
	"database/sql"
	"github.com/EricUCL/simplebank/api"
	db "github.com/EricUCL/simplebank/db/sqlc"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:123456@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "localhost:8080"
)
