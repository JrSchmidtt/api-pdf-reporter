package server

import (
	"log"

	"github.com/JrSchmidtt/api-gin/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer() Server{
	err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	return Server{
		port: "3000",
		server: gin.Default(),
	}
}

func (s *Server) Run(){
	router := routes.ConfigRoutes(s.server)
	log.Fatal(router.Run(":" + s.port))
}