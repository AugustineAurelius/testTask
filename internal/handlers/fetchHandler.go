package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/http"
	"test/internal/binance"
	"test/internal/model/httpmodel"
)

func FetchHandler(c *fiber.Ctx) error {
	request := new(httpmodel.FetchRequest)

	if err := c.BodyParser(request); err != nil {
		logrus.Errorf("Error while body parse in POST handler")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	fetchBinance, err := binance.FetchBinance(request)
	if err != nil {
		logrus.Errorf("Bad request while try to fetch")
		return err
	}
	return c.Send(fetchBinance)
}
