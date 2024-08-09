package routes

import (
	"github.com/ScoobieNoobie/Caterpillar/internal/controller"
	"github.com/ScoobieNoobie/Caterpillar/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func BatteryRoute(app *fiber.App) {
	app.Post("/create-battery", middleware.AuthenticateAndAuthorizeUser(), controller.InsertBattery)
	app.Post("/battery/images", middleware.AuthenticateAndAuthorizeUser(), controller.UploadBatteryImage)
	app.Get("/batter/:inspection_id", middleware.AuthenticateAndAuthorizeUser(), controller.GetBattery)
}
