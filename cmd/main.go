package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ScoobieNoobie/Caterpillar/config"
	"github.com/ScoobieNoobie/Caterpillar/internal/database"
	"github.com/ScoobieNoobie/Caterpillar/internal/routes"
	"github.com/ScoobieNoobie/Caterpillar/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load the details")
	}

	fmt.Println(config.ConfigData)

	err = utils.CreateCloudinaryClient()
	if err != nil {
		log.Fatalln("Failed to connect to cloudinary")
	}

	err = database.ConnectToDB()
	if err != nil {
		log.Fatalln("Failed to connect to mongo DB")
	}
	log.Println("Successfully connected to the database")

	ctx := context.Background()

	err = database.RunMigration(ctx)
	if err != nil {
		log.Fatalln("Failed to run the migrations " + err.Error())
	}
	log.Println("Migration has run successfully")

	routes.UserRoute(app)
	routes.HeaderRoute(app)
	routes.TireRoute(app)
	routes.BatteryRoute(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8080")
}
