package main

import (
	"github.com/gin-gonic/gin"

	"example.com/api/db"
	"example.com/api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
