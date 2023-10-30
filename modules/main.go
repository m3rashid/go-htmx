package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m3rashid/go-htmx/modules/home"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", home.HomeRoute)
}
