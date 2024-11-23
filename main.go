package main

import (
	"data-invetaris/config"
	"data-invetaris/database"
	"data-invetaris/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()

	app := fiber.New()

	if config.Config.AppMode == "development" {
		app.Use(logger.New(logger.Config{
			Format: "[FIBER] ${time} |${status}|${latency}|${method}| \"${path}\"\n",
		}))
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	
	routes.Routes(app)

	port := config.Config.AppPort
	log.Printf("🔥Fiber run in port %s", port)
	log.Fatal(app.Listen(":" + port))
}
