package webserver

type CryptoLastRateResponse string

type CryptoHistoryResponse struct {
	Offset  uint16           `json:"offset"`
	Limit   uint16           `json:"limit"`
	Total   uint16           `json:"total"`
	History *[]CryptoHistory `json:"history"`
}

type CryptoHistory struct {
	Timestamp int64   `json:"timestamp"`
	Datetime  string  `json:"datetime"`
	Value     float32 `json:"value"`
}

type FiatHistoryResponse struct {
	Total   uint16                        `json:"total"`
	History map[string]map[string]float32 ` json:"history"`
}

type CryptoFiatResponse map[string]float32
