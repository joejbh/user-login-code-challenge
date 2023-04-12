package handler

import (
	"net/mail"
	"user-login-code-challenge/service"

	"github.com/gofiber/fiber/v2"
)

func validate(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func Login(c *fiber.Ctx) error {
	type UserData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	input := new(service.LoginInput)

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Malformed request",
			})
	}

	if !validate(input.Email) {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Invalid email address",
			})
	}

	user, err := service.Authenticate(*input)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Your email/password combination did not work",
			})
	}

	userData := UserData{
		Username: user.Username,
		Email:    user.Email,
	}

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Success login",
			"data":    userData,
		})
}
