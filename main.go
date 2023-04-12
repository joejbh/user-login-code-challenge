package main

import (
	"log"

	"user-login-code-challenge/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	app.Use(cors.New())

	routes.Setup(app)

	log.Fatal(app.Listen(":" + viper.GetString("PORT")))
}
