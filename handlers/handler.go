package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Home renders the home view
func Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Loka Academy",
	})
}

// NoutFound renders the 404 view
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).Render("404", nil)
}
