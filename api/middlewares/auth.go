package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klimentru1986/go-event-booking/common/utils"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "empty token"})
		return
	}

	userId, err := utils.ValidateToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}

	ctx.Set("userId", userId)

	ctx.Next()
}
