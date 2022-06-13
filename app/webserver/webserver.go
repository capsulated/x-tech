package webserver

import (
	"context"
	"fmt"
	"github.com/capsulated/x-tech/config"
	"github.com/capsulated/x-tech/dbmsprovider"
	"github.com/gofiber/fiber/v2"
	"log"
)

type WebServing interface {
	Listen() error
	Shutdown() error
}

type WebServer struct {
	ctx          *context.Context
	app          *fiber.App
	addr         string
	dbmsProvider *dbmsprovider.DbmsProvider
}

func NewWebServer(ctx *context.Context, cfg *config.Server, db *dbmsprovider.DbmsProvider) *WebServer {
	log.Println("starting webserver...")
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	w := &WebServer{
		ctx:          ctx,
		addr:         addr,
		dbmsProvider: db,
	}

	app := fiber.New()

	app.Get("/api/btcusdt", w.cryptoCurrentValue)
	app.Post("/api/btcusdt", w.cryptoHistory)

	app.Get("/api/currencies", w.fiatCurrentValue)
	app.Post("/api/currencies", w.fiatHistory)

	app.Get("/api/latest", w.latestRate)
	//app.GET("/api/latest/{CHAR_CODE}", s.Hello)

	w.app = app
	log.Printf("webserver %s instance created\n", addr)
	return w
}

func (w *WebServer) Listen() error {
	return w.app.Listen(w.addr)
}

func (w *WebServer) Shutdown() error {
	return w.app.Shutdown()
}
