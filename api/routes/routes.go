package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/klimentru1986/go-event-booking/middlewares"
)

func SetupRoutes(server *gin.Engine) {

	v1 := server.Group("/api/v1")

	auth := v1.Group("/")
	auth.Use(middlewares.Authenticate)
	auth.POST("/events", createEvent)
	auth.PUT("/events/:id", updateEvent)
	auth.DELETE("/events/:id", deleteEvent)
	auth.POST("/events/:id/registration", registerForEvent)
	auth.DELETE("/events/:id/registration", cancelRegistration)

	v1.GET("/events", getEvents)
	v1.GET("/events/:id", getEventByID)

	v1.POST("/signup", signup)
	v1.POST("/login", login)
}
