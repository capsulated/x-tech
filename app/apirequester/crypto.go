package apirequester

import (
	"github.com/capsulated/x-tech/dbmsprovider"
	"log"
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
	rate, err := strconv.ParseFloat(
		c.Data.Last,
		64,
	)
	if err != nil {
		return nil, err
	}

	log.Println("rate", rate)
	log.Println("rate * 10000", int64(rate*10000))
	return &dbmsprovider.Rate{
		Time:         time.UnixMilli(c.Data.Time),
		TickerSource: "BTC",
		TickerTarget: "USDT",
		Rate:         int64(rate * 10000),
	}, nil
}
