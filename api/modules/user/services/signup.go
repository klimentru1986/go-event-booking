package services

import (
	"github.com/klimentru1986/go-event-booking/common/dto"
	"github.com/klimentru1986/go-event-booking/common/models"
)

func Signup(userDto *dto.CreateUserDto) (*models.User, error) {

	user := models.NewUser(userDto.Email, userDto.Password)

	err := user.Create()

	if err != nil {
		return nil, err
	}

	return user, nil
}
