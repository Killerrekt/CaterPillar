package routes

import (
	"github.com/ScoobieNoobie/Caterpillar/internal/controller"
	"github.com/ScoobieNoobie/Caterpillar/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func ExteriorRoute(app *fiber.App) {
	app.Post("/create-exterior", middleware.AuthenticateAndAuthorizeUser(), controller.CreateExterior)
	app.Post("/exterior-image/:inspection_id", middleware.AuthenticateAndAuthorizeUser(), controller.AddExteriorImage)
	app.Post("/exterior/image/:inspection_id", middleware.AuthenticateAndAuthorizeUser(), controller.AddExteriorImages)
	app.Get("/tire/:inspection_id", middleware.AuthenticateAndAuthorizeUser(), controller.GetExterior)
}
