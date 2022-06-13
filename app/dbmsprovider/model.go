package dbmsprovider

import (
	"time"
)

type Rate struct {
	Time         time.Time `db:"created_at"`
	TickerSource string    `db:"ticker_source"`
	TickerTarget string    `db:"ticker_target"`
	Rate         int64     `db:"rate"`
}
