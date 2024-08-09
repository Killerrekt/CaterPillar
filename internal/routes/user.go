package routes

import (
	"github.com/ScoobieNoobie/Caterpillar/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/signup", controller.SignUp)
}
