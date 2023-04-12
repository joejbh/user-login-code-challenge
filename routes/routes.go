package routes

import (
	"user-login-code-challenge/handler"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/login", handler.Login)
}
