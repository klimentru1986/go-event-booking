package models

import (
	"errors"

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
	query := "INSERT INTO users (email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hasedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(&u.Email, hasedPassword)

	if err != nil {
		return err
	}

	userId, err := res.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
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
