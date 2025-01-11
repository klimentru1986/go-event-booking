package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klimentru1986/go-event-booking/db"
	"github.com/klimentru1986/go-event-booking/models"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event

	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Fail to parse data"})
		return
	}

	// TODO change to real
	event.ID = 1
	event.UserID = 1

	event.Save()

	ctx.JSON(http.StatusCreated, event)
}
