package rotes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/klimentru1986/go-event-booking/models"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to load data"})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func getEventByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad id"})
		return
	}

	event, err := models.GetEventByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad id"})
		return
	}

	ctx.JSON(http.StatusOK, event)

}

func createEvent(ctx *gin.Context) {
	var event models.Event

	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Fail to parse data"})
		return
	}

	// TODO change to real
	event.UserID = 1

	err = event.Save()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Fail to save event"})
		return
	}

	ctx.JSON(http.StatusCreated, event)
}
