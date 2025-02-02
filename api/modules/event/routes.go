package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klimentru1986/go-event-booking/common/models"
	"github.com/klimentru1986/go-event-booking/middlewares"
	"github.com/klimentru1986/go-event-booking/modules/event/services"
)

func SetupEventRoutes(group *gin.RouterGroup) {

	auth := group.Group("/")
	auth.Use(middlewares.Authenticate)
	auth.POST("/events", createEvent)
	auth.PUT("/events/:id", updateEvent)
	auth.DELETE("/events/:id", deleteEvent)
	auth.POST("/events/:id/registration", registerForEvent)
	auth.DELETE("/events/:id/registration", cancelRegistration)

	group.GET("/events", getEvents)
	group.GET("/events/:id", getEventByID)
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func getEventByID(ctx *gin.Context) {
	_, event, err := services.FindEventByStrId(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, event)

}

func createEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	var event models.Event

	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	err = services.CreateEvent(&event, userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusCreated, event)
}

func updateEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	var updatedEvent models.Event

	err := ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	err = services.UpdateEvent(ctx.Param("id"), &updatedEvent, userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, updatedEvent)
}

func deleteEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	err := services.DeleteEvent(ctx.Param("id"), userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	err := services.RegisterForEvent(ctx.Param("id"), userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

func cancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	err := services.CancelRegistration(ctx.Param("id"), userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
