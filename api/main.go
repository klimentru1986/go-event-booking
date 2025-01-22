package main

import (
	"github.com/gin-gonic/gin"
	"github.com/klimentru1986/go-event-booking/common/db"
	"github.com/klimentru1986/go-event-booking/modules/event"
	"github.com/klimentru1986/go-event-booking/modules/user"
)

func main() {
	// db.InitDB("sqlite3", "common/db/api.db")
	db.InitDB("postgres", "postgresql://goeventbooking:pass@localhost:5432/goeventbooking?sslmode=disable")
	server := gin.Default()

	v1 := server.Group("/api/v1")

	user.SetupUserRoutes(v1)
	event.SetupEventRoutes(v1)

	server.Run(":8080")
}
