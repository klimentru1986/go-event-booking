package main

import (
	"github.com/gin-gonic/gin"
	"github.com/klimentru1986/go-event-booking/common/config"
	"github.com/klimentru1986/go-event-booking/common/db"
	"github.com/klimentru1986/go-event-booking/modules/event"
	"github.com/klimentru1986/go-event-booking/modules/user"
)

func main() {
	conf := config.New()

	db.InitDB(conf.DbDriver, conf.DbSource)
	server := gin.Default()

	v1 := server.Group("/api/v1")

	user.SetupUserRoutes(v1)
	event.SetupEventRoutes(v1)

	server.Run(":8080")
}
