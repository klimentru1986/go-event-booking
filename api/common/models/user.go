package models

import (
	"errors"
	"fmt"

	"github.com/klimentru1986/go-event-booking/common/db"
	"github.com/klimentru1986/go-event-booking/common/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required" `
	Password string `json:"-" binding:"required" `
}

func NewUser(email string, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

func (u *User) Create() error {
	var userId int64

	hasedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO users (email, password) VALUES ('%s' , '%s') RETURNING id", u.Email, hasedPassword)

	err = db.DB.QueryRow(query).Scan(&userId)

	if err != nil {
		return err
	}

	u.ID = userId
	return err
}
func (u *User) Delete() error {
	query := "DELETE FROM users WHERE email = $1"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Email)

	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = $1"
	row := db.DB.QueryRow(query, u.Email)

	var hashedPassword string
	err := row.Scan(&u.ID, &hashedPassword)
	if err != nil {
		return err
	}

	isValid := utils.ComparePassword(hashedPassword, u.Password)

	if !isValid {
		return errors.New("invalid password")
	}

	return nil
}
