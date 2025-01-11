package routes

import (
	"net/http"

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
	_, event, err := findEventByParam(ctx.Param("id"))

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

	err = event.Create()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Fail to save event"})
		return
	}

	ctx.JSON(http.StatusCreated, event)
}

func updateEvent(ctx *gin.Context) {
	id, _, err := findEventByParam(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad id"})
		return
	}

	var updatedEvent models.Event

	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Fail to parse data"})
		return
	}

	updatedEvent.ID = *id

	err = updatedEvent.Update()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to update data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func deleteEvent(ctx *gin.Context) {
	_, event, err := findEventByParam(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad id"})
		return
	}

	err = event.Delete()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
