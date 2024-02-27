package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/mahdikarami0111/simplebank/api"
	db "github.com/mahdikarami0111/simplebank/db/sqlc"
	"github.com/mahdikarami0111/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cant load config")
	}
	dbConn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cant connect to db", err)
	}

	store := db.NewStore(dbConn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cant start server", err)
	}

}
