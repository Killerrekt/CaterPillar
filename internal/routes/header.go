package routes

import (
	"github.com/ScoobieNoobie/Caterpillar/internal/controller"
	"github.com/ScoobieNoobie/Caterpillar/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func HeaderRoute(app *fiber.App) {
	app.Post("/create-header", middleware.AuthenticateAndAuthorizeUser(), controller.UploadHeader)
	app.Get("/get-header", middleware.AuthenticateAndAuthorizeUser(), controller.GetHeader)
}
