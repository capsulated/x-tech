package apirequester

import (
	"github.com/capsulated/x-tech/dbmsprovider"
	"strconv"
	"time"
)

type CryptoResponse struct {
	Code string `json:"code"`
	Data struct {
		Time             int64  `json:"time"`
		Symbol           string `json:"symbol"`
		Buy              string `json:"buy"`
		Sell             string `json:"sell"`
		ChangeRate       string `json:"changeRate"`
		ChangePrice      string `json:"changePrice"`
		High             string `json:"high"`
		Low              string `json:"low"`
		Vol              string `json:"vol"`
		VolValue         string `json:"volValue"`
		Last             string `json:"last"`
		AveragePrice     string `json:"averagePrice"`
		TakerFeeRate     string `json:"takerFeeRate"`
		MakerFeeRate     string `json:"makerFeeRate"`
		TakerCoefficient string `json:"takerCoefficient"`
		MakerCoefficient string `json:"makerCoefficient"`
	} `json:"data"`
}

func (c *CryptoResponse) ToRate() (*dbmsprovider.Rate, error) {
	rate, err := strconv.ParseFloat(c.Data.Last, 32)
	if err != nil {
		return nil, err
	}

	return &dbmsprovider.Rate{
		Time:     time.UnixMilli(c.Data.Time),
		Base:     "BTC",
		Currency: "USDT",
		Rate:     float32(rate),
	}, nil
}
