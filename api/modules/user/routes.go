package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klimentru1986/go-event-booking/common/dto"
	"github.com/klimentru1986/go-event-booking/modules/user/services"
)

func SetupUserRoutes(group *gin.RouterGroup) {

	group.POST("/signup", signup)
	group.POST("/login", login)
}

func signup(ctx *gin.Context) {
	var userDto dto.CreateUserDto
	err := ctx.BindJSON(&userDto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	user, err := services.Signup(&userDto)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func login(ctx *gin.Context) {
	var userDto dto.CreateUserDto
	err := ctx.BindJSON(&userDto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	jwt, err := services.Login(&userDto)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"token": jwt})
}
