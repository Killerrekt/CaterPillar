package routes

import (
	"github.com/ScoobieNoobie/Caterpillar/internal/controller"
	"github.com/ScoobieNoobie/Caterpillar/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func BrakeRoute(app *fiber.App) {
	app.Post("/create-brake", middleware.AuthenticateAndAuthorizeUser(), controller.CreateBrake)
	app.Post("/brake/image/:inspection_id", middleware.AuthenticateAndAuthorizeUser(), controller.AddBrakeImages)
	app.Get("/tire/:inspection_id", middleware.AuthenticateAndAuthorizeUser(), controller.GetBrake)
}
