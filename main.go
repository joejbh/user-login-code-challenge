package main

import (
	"log"

	"user-login-code-challenge/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] | ${pid} | ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
	}))

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Use(cors.New())

	routes.Setup(app)

	log.Fatal(app.Listen(":" + viper.GetString("PORT")))
}
