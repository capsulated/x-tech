package worker

import (
	"github.com/capsulated/x-tech/apirequester"
	"github.com/capsulated/x-tech/config"
	"github.com/capsulated/x-tech/dbmsprovider"
	"github.com/robfig/cron/v3"
	"log"
)

type Working interface {
	Start()
	fiatWork()
	cryptoWork()
	Stop()
}

type Worker struct {
	cron         *cron.Cron
	cryptoCronId *cron.EntryID
	fiatCronId   *cron.EntryID
	apiRequester *apirequester.ApiRequester
	dbmsProvider *dbmsprovider.DbmsProvider
}

func NewWorker(cfg *config.Cron, api *apirequester.ApiRequester, db *dbmsprovider.DbmsProvider) *Worker {
	w := &Worker{
		cron:         cron.New(),
		apiRequester: api,
		dbmsProvider: db,
	}

	cid, err := w.cron.AddFunc(cfg.Crypto, w.cryptoWork)
	if err != nil {
		log.Fatalln(err)
	}

	fid, err := w.cron.AddFunc(cfg.Fiat, w.fiatWork)
	if err != nil {
		log.Fatalln(err)
	}

	w.cryptoCronId = &cid
	w.fiatCronId = &fid

	return w
}

func (w *Worker) Start() {
	w.fiatWork()
	w.cryptoWork()

	w.cron.Start()
}

func (w *Worker) cryptoWork() {
	log.Println("crypto work...")
	rate, err := w.apiRequester.CryptoRequest()
	if err != nil {
		return
	}

	if w.dbmsProvider.LastCrypto != nil && rate.Time == w.dbmsProvider.LastCrypto.Time {
		log.Println("crypto last updated time the same, return")
		return
	}

	err = w.dbmsProvider.InsertCryptoRate(rate)
	if err != nil {
		return
	}
	w.dbmsProvider.LastCrypto = rate

	log.Println("crypto work success!")
}

func (w *Worker) fiatWork() {
	log.Println("fiat work...")
	rates, t, err := w.apiRequester.FiatRequest()
	if err != nil {
		return
	}

	if w.dbmsProvider.LastFiat.Time != nil && *t == *w.dbmsProvider.LastFiat.Time {
		log.Println("fiat last updated time the same, return")
		return
	}

	for _, rate := range *rates {
		err = w.dbmsProvider.InsertFiatRate(&rate)
		if err != nil {
			return
		}
	}
	w.dbmsProvider.LastFiat.Rates = rates
	w.dbmsProvider.LastFiat.Time = t

	log.Println("fiat work success!")
}

func (w *Worker) Stop() {
	w.cron.Remove(*w.fiatCronId)
	w.cron.Remove(*w.cryptoCronId)
}
