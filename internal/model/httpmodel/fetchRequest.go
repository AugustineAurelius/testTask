package httpmodel

type FetchRequest struct {
	Ticker   string `json:"ticker"`
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`
}
