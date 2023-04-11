package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
    app := fiber.New()
	viper.SetConfigFile(".env")
    viper.ReadInConfig()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    log.Fatal(app.Listen(":" + viper.GetString("PORT")))
}