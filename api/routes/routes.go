package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/klimentru1986/go-event-booking/middlewares"
)

func SetupRoutes(group *gin.RouterGroup) {

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
