package dbmsprovider

import (
	"context"
	"fmt"
	"github.com/capsulated/x-tech/config"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type DbmsProviding interface {
	InsertCryptoRate(*Rate) error
	InsertFiatRate(*Rate) error
	SelectRangeCryptoRates(string, string, string, string) (*[]Rate, error)
	SelectRangeFiatRates(string, string, string, string) (*[]Rate, error)
	Close() error
}

type DbmsProvider struct {
	ctx        *context.Context
	Db         *sqlx.DB
	LastCrypto *Rate
	LastFiat   struct {
		Rates *[]Rate
		Time  *time.Time
	}
}

func NewDbmsProvider(ctx *context.Context, cfg *config.Dbms) (dbp *DbmsProvider) {
	connStr := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		cfg.Host, cfg.Port, cfg.DbName, cfg.User, cfg.Password,
	)

	db, err := sqlx.ConnectContext(*ctx, "pgx", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	db.DB.SetMaxIdleConns(10)

	dbp = &DbmsProvider{
		ctx: ctx,
		Db:  db,
	}
	return
}

func (p *DbmsProvider) Close() error {
	return p.Db.Close()
}
