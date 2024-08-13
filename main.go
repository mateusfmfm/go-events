package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusfmfm/go-events/db"
	"github.com/mateusfmfm/go-events/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
