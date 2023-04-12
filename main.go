package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	app.Use(cors.New())

	app.Post("/api/login", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"foo": "bar",
		})
	})

	log.Fatal(app.Listen(":" + viper.GetString("PORT")))
}
