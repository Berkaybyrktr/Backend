package main

import (
	"database/sql"
	"github/techschool/simplebank/Api"
	"github/techschool/simplebank/Util"

	db "github/techschool/simplebank/db/sqlc"
	"log"
)

func main() {
	config, err := Util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the db", err)
	}
	store := db.NewStore(conn)
	server := Api.NewServer(store)

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start the server", err)
	}
}
