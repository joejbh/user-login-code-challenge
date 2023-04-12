package service

import (
	"errors"
	"time"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Username string
	Email    string
}

func Authenticate(loginInput LoginInput) (*User, error) {
	time.Sleep(1500 * time.Millisecond)

	if loginInput.Email != "test@test.com" {
		return nil, errors.New("record not found")
	}

	if loginInput.Password != "Password1" {
		return nil, errors.New("bad password")
	}

	user := User{
		Username: "Test User",
		Email:    "test@test.com",
	}

	return &user, nil
}
