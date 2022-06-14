package apirequester

import (
	"encoding/xml"
	"github.com/capsulated/x-tech/dbmsprovider"
	"strconv"
	"strings"
	"time"
)

type FiatResponse struct {
	XMLName xml.Name `xml:"ValCurs"`
	Text    string   `xml:",chardata"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valute  []struct {
		Text     string `xml:",chardata"`
		ID       string `xml:"ID,attr"`
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Nominal  string `xml:"Nominal"`
		Name     string `xml:"Name"`
		Value    string `xml:"Value"`
	} `xml:"Valute"`
}

func (f *FiatResponse) ToFiatRates() (*[]dbmsprovider.Rate, *time.Time, error) {
	t, err := time.Parse("02.01.2006", f.Date)
	if err != nil {
		return nil, nil, err
	}

	var rates []dbmsprovider.Rate
	for _, v := range f.Valute {
		rate, err := strconv.ParseFloat(
			strings.Replace(v.Value, ",", ".", 1), 64,
		)
		if err != nil {
			return nil, nil, err
		}

		nominal, err := strconv.ParseFloat(v.Nominal, 64)
		if err != nil {
			return nil, nil, err
		}

		rates = append(rates, dbmsprovider.Rate{
			Time:     t,
			Base:     v.CharCode,
			Currency: "RUB",
			Rate:     float32(rate / nominal),
		})
	}

	rates = append(rates, dbmsprovider.Rate{
		Time:     t,
		Base:     "RUB",
		Currency: "RUB",
		Rate:     1,
	})

	return &rates, &t, nil
}
