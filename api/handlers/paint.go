package handlers

import (
	"digitalrepublic/pkg/paint"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func PaintSizes() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var requestBody paint.CalculateRoomPaintInCansInput

		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(&fiber.Map{
				"Error": "Valores dos campos invalidos, confira os campos e tente novamente",
			})
		}

		interactor := paint.NewCalculateRoomPaintInCans()
		result, err := interactor.Execute(requestBody)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"Error": err.Error(),
			})

		}
		return c.JSON(result)

	}

}
