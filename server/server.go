package server

import (
	"books-api/server/routes"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port: "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Print("server is running in port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}