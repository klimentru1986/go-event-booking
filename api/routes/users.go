package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klimentru1986/go-event-booking/dto"
	"github.com/klimentru1986/go-event-booking/models"
)

func signup(ctx *gin.Context) {
	var userDto dto.CreateUserDto
	err := ctx.BindJSON(&userDto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	user := models.NewUser(userDto.Email, userDto.Password)

	err = user.Create()

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

	user := models.NewUser(userDto.Email, userDto.Password)

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
