package binance

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
	"strings"
	"test/internal/model/httpmodel"
	"time"
)

func FetchBinance(request *httpmodel.FetchRequest) ([]byte, error) {
	klineOpen := timeParser(request.DateFrom)
	klineClose := timeParser(request.DateTo)
	tickerForResponse := strings.ReplaceAll(request.Ticker, "/", "")

	responseFromBinance, err := http.Get(
		fmt.Sprintf("https://api.binance.com/api/v3/klines?symbol=%s&interval=1m&startTime=%s&endTime=%s",
			tickerForResponse, strconv.Itoa(int(klineOpen)), strconv.Itoa(int(klineClose))),
	)
	if err != nil {
		logrus.Fatalln("Error while send get request in binance API", err)
	}

	b, err := io.ReadAll(responseFromBinance.Body)
	if err != nil {
		logrus.Error("Couldn't close request body")
	}
	BinanceString := string(b)
	binanceArray := strings.Split(BinanceString, ",")

	startValue, err := strconv.ParseFloat(strings.ReplaceAll(binanceArray[2], "\"", ""), 32)
	if err != nil {
		logrus.Error("Error while parse float")
	}

	endValue, err := strconv.ParseFloat(strings.ReplaceAll(binanceArray[len(binanceArray)-8], "\"", ""), 32)
	if err != nil {
		logrus.Error("Error while parse float")
	}
	difference := ((endValue - startValue) / startValue) * 100

	response := httpmodel.FetchResponse{Ticker: request.Ticker,
		Price:      endValue,
		Difference: fmt.Sprintf("%f`%", difference),
	}
	marshal, err := json.Marshal(response)
	if err != nil {
		logrus.Error("Error while marshal json")
		return nil, err
	}
	return marshal, nil

}

func timeParser(dateString string) int {
	dateFormat := "02.01.2006 15:04:05"

	dt, err := time.Parse(dateFormat, dateString)
	if err != nil {
		fmt.Println("Ошибка при парсинге даты и времени:", err)
		return 0
	}
	return int(dt.UnixNano() / 1_000_000)
}
