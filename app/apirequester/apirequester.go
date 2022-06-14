package apirequester

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/capsulated/x-tech/config"
	"github.com/capsulated/x-tech/dbmsprovider"
	"github.com/mailru/easyjson"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ApiRequesting interface {
	CryptoRequest() (*CryptoResponse, error)
	FiatRequest() (*FiatResponse, error)
}

type ApiRequester struct {
	ctx             *context.Context
	cryptoSourceUrl string
	fiatSourceUrl   string
}

func NewApiRequester(ctx *context.Context, c *config.Api) *ApiRequester {
	return &ApiRequester{
		ctx,
		c.CryptoSourceUrl,
		c.FiatSourceUrl,
	}
}

func (a *ApiRequester) CryptoRequest() (r *dbmsprovider.Rate, err error) {
	req, err := http.NewRequestWithContext(*a.ctx, http.MethodGet, a.cryptoSourceUrl, nil)
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("err:", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err:", err)
		return
	}

	c := &CryptoResponse{}
	err = easyjson.Unmarshal(body, c)
	if err != nil {
		log.Println("err:", err)
		return
	}

	r, err = c.ToRate()
	if err != nil {
		log.Println("err:", err)
		return
	}

	return
}

func (a *ApiRequester) FiatRequest() (*[]dbmsprovider.Rate, *time.Time, error) {
	req, err := http.NewRequestWithContext(*a.ctx, http.MethodGet, a.fiatSourceUrl, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	fiatResponse := &FiatResponse{}
	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&fiatResponse)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	rates, ratesTime, err := fiatResponse.ToFiatRates()
	if err != nil {
		log.Println("err:", err)
		return nil, nil, err
	}

	return rates, ratesTime, nil
}
