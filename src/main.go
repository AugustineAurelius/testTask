package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"log/slog"
	"test/internal/binance"
	"test/internal/config"
	"test/internal/connections"
	"test/internal/handlers"
	"test/internal/model/databasemodel"
	"time"
)

var router = fiber.New()
var cfg config.Config
var logger *logrus.Logger

func init() {

	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		slog.Info("Could not load environment variables", err)
	}
	connections.ConnectDB(&cfg)

	logger = logrus.New()

	router.Post("/add_ticker", handlers.PostHandler)
	router.Get("/fetch", handlers.FetchHandler)

}

func main() {
	err := connections.DB.AutoMigrate(&databasemodel.Ticker{}, &databasemodel.TickerWithPrice{})
	if err != nil {
		logger.Error("Could not migrate database", err)
	}

	go func() {

		err = router.Listen(cfg.APIPort) // инизиализируем сервер
		if err != nil {
			logger.Fatal("Could not to start listen on port")
		}
	}()

	for {
		time.Sleep(1 * time.Minute)
		go binance.GetPriceFromBinance(connections.DB)
	}

}
