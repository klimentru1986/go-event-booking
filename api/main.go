package main

import (
	"github.com/gin-gonic/gin"
	"github.com/klimentru1986/go-event-booking/db"
	"github.com/klimentru1986/go-event-booking/rotes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	rotes.SetupRoutes(server)

	server.Run(":8080")
}
