package home

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func HomeRoute(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": os.Getenv("APP_NAME"),
		"headerItems": fiber.Map{
			"Dashboard": "/dashboard",
			"Team":      "/team",
			"Projects":  "/projects",
			"Calender":  "/calender",
		},
		"activeHeader": "/",
	})
}
