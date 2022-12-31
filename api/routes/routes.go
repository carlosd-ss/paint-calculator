package routes

import (
	"digitalrepublic/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router) {
	app.Get("/amount-of-paint", handlers.PaintSizes())
}
