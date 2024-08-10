package routes

import (
	"github.com/ScoobieNoobie/Caterpillar/internal/controller"
	"github.com/ScoobieNoobie/Caterpillar/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func EngineRoute(app *fiber.App) {
	app.Post("/create-engine", middleware.AuthenticateAndAuthorizeUser(), controller.CreateExterior)
	app.Post("/engine-image/:inspection_id", middleware.AuthenticateAndAuthorizeUser(), controller.AddExteriorImage)
	app.Post("/engine/image/:inspection_id", middleware.AuthenticateAndAuthorizeUser(), controller.AddExteriorImages)
	app.Get("/engine/:inspection_id", middleware.AuthenticateAndAuthorizeUser(), controller.GetExterior)
}
