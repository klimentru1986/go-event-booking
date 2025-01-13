package services

import (
	"github.com/klimentru1986/go-event-booking/common/dto"
	"github.com/klimentru1986/go-event-booking/common/models"
	"github.com/klimentru1986/go-event-booking/common/utils"
)

func Login(userDto *dto.CreateUserDto) (string, error) {

	user := models.NewUser(userDto.Email, userDto.Password)

	err := user.ValidateCredentials()

	if err != nil {
		return "", err
	}

	return utils.GenerateToken(user.ID, user.Email)
}
