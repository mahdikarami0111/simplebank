package api

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	db "github.com/mahdikarami0111/simplebank/db/sqlc"
	"github.com/mahdikarami0111/simplebank/util"
)

var server *Server

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	config, err := util.LoadConfig("../.")
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cant connect to db", err)
	}

	store := db.NewStore(dbConn)
	server = NewServer(store)

	os.Exit(m.Run())
}
