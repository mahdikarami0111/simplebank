package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/mahdikarami0111/simplebank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//rout shit
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.POST("/transfers", server.CreateTransfer)
	router.POST("/users", server.createUser)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	fmt.Println("Server starting ...")
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
