package webserver

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func (w *WebServer) cryptoCurrentValue(c *fiber.Ctx) error {
	return c.JSON(w.dbmsProvider.LastCrypto)
}

func (w *WebServer) cryptoHistory(c *fiber.Ctx) error {
	start, err := strconv.ParseInt(c.Query("start"), 10, 16)
	end, err := strconv.ParseInt(c.Query("end"), 10, 16)
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 16)
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 16)

	if err != nil || c.Query("start") == "" || c.Query("end") == "" || c.Query("offset") == "" || c.Query("limit") == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	rates, total, err := w.dbmsProvider.SelectRangeCryptoRates(int16(start), int16(end), int16(offset), int16(limit))
	if err != nil {
		log.Printf("dbmsProvider.SelectRangeCryptoRates() start %d, end %d, offset %d, limit %d err: %s",
			start, end, offset, limit, err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var history []CryptoHistory
	for _, r := range *rates {
		timestamp := r.Time.Unix()
		value := float64(r.Rate) / 1000
		history = append(history, CryptoHistory{timestamp, value})
	}

	resp := &CryptoHistoryResponse{
		Offset:  int16(offset),
		Limit:   int16(offset),
		Total:   total,
		History: &history,
	}

	return c.JSON(resp)
}

func (w *WebServer) fiatCurrentValue(c *fiber.Ctx) error {
	return c.JSON(w.dbmsProvider.LastFiat)
}

func (w *WebServer) fiatHistory(c *fiber.Ctx) error {
	start := c.Query("start")
	end := c.Query("end")
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 16)
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 16)

	if err != nil || start == "" || end == "" || c.Query("offset") == "" || c.Query("limit") == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	_, total, err := w.dbmsProvider.SelectRangeFiatRates(start, end, int16(offset), int16(limit))
	if err != nil {
		log.Printf("dbmsProvider.SelectRangeCryptoRates() start %s, end %s, offset %d, limit %d err: %s",
			start, end, offset, limit, err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var history []FiatHistory

	// Sort for FiatHistory with time and FiatRate
	//for _, r := range *rates {
	//	timestamp := r.Time.Unix()
	//	FiatRat
	//	history = append(history, FiatHistory{timestamp, value})
	//}

	resp := &FiatHistoryResponse{
		Offset:  int16(offset),
		Limit:   int16(offset),
		Total:   total,
		History: &history,
	}
	// todo get last fiat value by time range

	return c.JSON(resp)
}

func (*WebServer) latestRate(c *fiber.Ctx) error {
	return c.SendString("latestRate")
}
