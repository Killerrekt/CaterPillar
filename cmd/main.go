package main

import (
	"fmt"
	"log"

	"github.com/ScoobieNoobie/Caterpillar/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load the details")
	}

	fmt.Println(config.ConfigData)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8080")
}
