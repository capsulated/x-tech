package main

import (
	"context"
	"github.com/capsulated/x-tech/apirequester"
	"github.com/capsulated/x-tech/config"
	"github.com/capsulated/x-tech/dbmsprovider"
	"github.com/capsulated/x-tech/webserver"
	"github.com/capsulated/x-tech/worker"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := config.NewConfig()

	ctx, cancel := context.WithCancel(context.Background())

	api := apirequester.NewApiRequester(&ctx, &cfg.Api)

	db := dbmsprovider.NewDbmsProvider(&ctx, &cfg.Dbms)

	srv := webserver.NewWebServer(&ctx, &cfg.Server, db)
	go func() {
		err := srv.Listen()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	wrk := worker.NewWorker(&cfg.Cron, api, db)
	go wrk.Start()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT)
	for {
		select {
		case sig := <-signals:
			log.Printf("OS cmd received signal %s", sig.String())
			switch sig {
			case syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT:
				cancel()
				err := srv.Shutdown()
				if err != nil {
					log.Println("err:", err)
				}
				err = db.Close()
				if err != nil {
					log.Println("err:", err)
				}
				wrk.Stop()
				log.Println("application shutdown gracefully")
				os.Exit(0)
			case syscall.SIGABRT:
				os.Exit(1)
			}
			break
		}
	}
}
