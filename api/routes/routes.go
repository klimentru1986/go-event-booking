package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(server *gin.Engine) {

	v1 := server.Group("/api/v1")

	v1.GET("/events", getEvents)
	v1.GET("/events/:id", getEventByID)
	v1.POST("/events", createEvent)
	v1.PUT("/events/:id", updateEvent)
	v1.DELETE("/events/:id", deleteEvent)

	v1.POST("/signup", signup)
}
