package main

import (
	"github.com/gin-gonic/gin"
	"github.com/klimentru1986/go-event-booking/common/db"
	"github.com/klimentru1986/go-event-booking/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.SetupRoutes(server)

	server.Run(":8080")
}
