package dbmsprovider

import (
	"time"
)

type Rate struct {
	Time     time.Time `db:"time"`
	Currency string    `db:"currency"`
	Base     string    `db:"base"`
	Rate     float32   `db:"rate"`
}
