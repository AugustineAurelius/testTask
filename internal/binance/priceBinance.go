package binance

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strconv"
	"strings"
	"test/internal/model/databasemodel"
	"test/internal/model/httpmodel"
)

func GetPriceFromBinance(db *gorm.DB) {

	var listOfTickers []databasemodel.Ticker
	db.Find(&listOfTickers)

	for i := 0; i < len(listOfTickers); i++ {
		currentTicker := listOfTickers[i]

		tickerForRequest := strings.ReplaceAll(currentTicker.Ticker, "/", "")
		response, err := http.Get(
			fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", tickerForRequest),
		)
		if err != nil {
			logrus.Fatalln("Error while send get request in binance API", err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				logrus.Error("Couldn't close request body")
			}
		}(response.Body)

		b, err := io.ReadAll(response.Body)
		if err != nil {
			logrus.Error("Error while read all body")
		}

		var pr httpmodel.PriceResponse

		err = json.Unmarshal(b, &pr)
		if err != nil {
			logrus.Error("Couldn't unmarshall json")
		}
		tickerPrice, err := strconv.ParseFloat(pr.Price, 64)
		if err != nil {
			logrus.Error("Error while parse float")
		}
		tickerForDB := databasemodel.TickerWithPrice{
			Price:  tickerPrice,
			Ticker: currentTicker.Ticker,
		}
		db.Create(&tickerForDB)
		logrus.Info("Successfully create new ticker with price", tickerForDB)
	}
}
