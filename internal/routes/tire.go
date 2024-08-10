package routes

import (
	"github.com/ScoobieNoobie/Caterpillar/internal/controller"
	"github.com/ScoobieNoobie/Caterpillar/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func TireRoute(app *fiber.App) {
	app.Post("/create-tire", middleware.AuthenticateAndAuthorizeUser(), controller.InsertTire)
	app.Post("/tire/images/:inspection_id", middleware.AuthenticateAndAuthorizeUser(), controller.UploadTireImage)
	app.Get("/tire/:inspection_id", middleware.AuthenticateAndAuthorizeUser(), controller.GetTire)
	app.Get("/tire/images/:inspection_id", middleware.AuthenticateAndAuthorizeUser(), controller.GetTireImages)
}
