package webserver

type CryptoLastRateResponse string

type CryptoHistoryResponse struct {
	Offset  int16            `json:"offset"`
	Limit   int16            `json:"limit"`
	Total   int16            `json:"total"`
	History *[]CryptoHistory `json:"history"`
}

type CryptoHistory struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

type FiatHistoryResponse struct {
	Offset  int16          `json:"offset"`
	Limit   int16          `json:"limit"`
	Total   int16          `json:"total"`
	History *[]FiatHistory `json:"history"`
}

type FiatHistory struct {
	Date     string `json:"date"`
	FiatRate *FiatRate
}

type FiatRate map[string]float64

type CryptoFiatResponse map[string]float64
