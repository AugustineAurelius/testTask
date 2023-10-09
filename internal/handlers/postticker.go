package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/http"
	"test/internal/connections"
	"test/internal/model/databasemodel"
)

func PostHandler(c *fiber.Ctx) error {
	ticker := new(databasemodel.Ticker)

	if err := c.BodyParser(ticker); err != nil {
		logrus.Errorf("Error while body parse in POST handler")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := connections.DB.Create(ticker).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(ticker)
}
