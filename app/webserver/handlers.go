package webserver

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func (w *WebServer) cryptoCurrentValue(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("%g", w.dbmsProvider.LastCrypto.Rate))
}

func (w *WebServer) cryptoHistory(c *fiber.Ctx) error {
	start := c.Query("start")
	end := c.Query("end")
	offset := c.Query("offset")
	limit := c.Query("limit")

	if start == "" || end == "" || offset == "" || limit == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	rates, total, err := w.dbmsProvider.SelectRangeCryptoRates(start, end, offset, limit)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var history []CryptoHistory
	for _, r := range *rates {
		timestamp := r.Time.Unix()
		datetime := r.Time.Format("2006-01-02 15:04:05")
		history = append(history, CryptoHistory{timestamp, datetime, r.Rate})
	}

	o, err := strconv.ParseUint(offset, 10, 16)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	l, err := strconv.ParseUint(limit, 10, 16)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	resp := &CryptoHistoryResponse{
		Offset:  uint16(o),
		Limit:   uint16(l),
		Total:   total,
		History: &history,
	}

	return c.JSON(resp)
}

func (w *WebServer) fiatCurrentValue(c *fiber.Ctx) error {
	fiatRate := make(map[string]float32)
	log.Printf("%#v", w.dbmsProvider.LastFiat)
	for _, r := range *w.dbmsProvider.LastFiat.Rates {
		fiatRate[r.Base] = r.Rate
	}
	return c.JSON(fiatRate)
}

func (w *WebServer) fiatHistory(c *fiber.Ctx) error {
	start := c.Query("start")
	end := c.Query("end")

	if start == "" || end == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	rates, err := w.dbmsProvider.SelectRangeFiatRates(start, end)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	fiatHistoryResponse := &FiatHistoryResponse{
		History: make(map[string]map[string]float32),
	}
	for _, rate := range *rates {
		date := rate.Time.Format("2006-01-02")
		if _, ok := fiatHistoryResponse.History[date]; !ok {
			fiatHistoryResponse.History[date] = make(map[string]float32)
		}
		fiatHistoryResponse.History[date][rate.Base] = rate.Rate
	}
	fiatHistoryResponse.Total = uint16(len(fiatHistoryResponse.History))

	return c.JSON(fiatHistoryResponse)
}

func (w *WebServer) latestRate(c *fiber.Ctx) error {
	cryptoFiatResponse := make(CryptoFiatResponse)

	for _, r := range *w.dbmsProvider.LastFiat.Rates {
		cryptoFiatResponse[r.Base] = w.dbmsProvider.LastCrypto.Rate * (*w.dbmsProvider.LastFiat.UsdRub / r.Rate)
	}

	return c.JSON(cryptoFiatResponse)
}
