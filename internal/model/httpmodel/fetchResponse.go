package httpmodel

type FetchResponse struct {
	Ticker     string  `json:"ticker"`
	Price      float64 `json:"price"`
	Difference string  `json:"difference"`
}
