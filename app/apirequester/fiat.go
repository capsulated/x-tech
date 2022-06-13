package apirequester

import (
	"encoding/xml"
	"fmt"
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

func (f *FiatResponse) ToRates() (*[]dbmsprovider.Rate, *time.Time, error) {
	t, err := time.Parse("02.01.2006", f.Date)
	if err != nil {
		return nil, nil, err
	}

	var rates []dbmsprovider.Rate
	for _, v := range f.Valute {
		var rate int64
		rate, err = strconv.ParseInt(
			fmt.Sprintf("%s000", strings.Replace(v.Value, ",", "", 1)),
			10,
			64,
		)
		if err != nil {
			return nil, nil, err
		}

		rates = append(rates, dbmsprovider.Rate{
			Time:         t,
			TickerSource: v.CharCode,
			TickerTarget: "RUB",
			Rate:         rate,
		})
	}

	return &rates, &t, nil
}
