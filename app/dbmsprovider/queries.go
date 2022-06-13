package dbmsprovider

import "log"

func (p *DbmsProvider) InsertCryptoRate(r *Rate) (err error) {
	log.Printf("save crypto rate %s-%s %d...\n", r.TickerSource, r.TickerTarget, r.Rate)
	res, err := p.Db.Exec(
		"insert into crypto (time, ticker_source, ticker_target, rate) values ($1, $2, $3, $4) on conflict do nothing",
		r.Time,
		r.TickerSource,
		r.TickerTarget,
		r.Rate,
	)
	if err != nil {
		log.Printf("save fiat crypto %s-%s %d err: %s\n", r.TickerSource, r.TickerTarget, r.Rate, err)
		return
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		log.Printf("crypto rate %s-%s %d already exist\n", r.TickerSource, r.TickerTarget, r.Rate)
		return
	}

	p.LastCrypto = r
	log.Printf("crypto rate %s-%s %d saved!\n", r.TickerSource, r.TickerTarget, r.Rate)
	return
}

func (p *DbmsProvider) InsertFiatRate(r *Rate) (err error) {
	log.Printf("save fiat rate %s-%s %d...\n", r.TickerSource, r.TickerTarget, r.Rate)
	res, err := p.Db.Exec(
		"insert into fiat (time, ticker_source, ticker_target, rate) VALUES ($1, $2, $3, $4) on conflict do nothing",
		r.Time,
		r.TickerSource,
		r.TickerTarget,
		r.Rate,
	)
	if err != nil {
		log.Printf("save fiat rate %s-%s %d err: %s\n", r.TickerSource, r.TickerTarget, r.Rate, err)
		return
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		log.Printf("fiat rate %s-%s %d already exist\n", r.TickerSource, r.TickerTarget, r.Rate)
		return
	}

	log.Printf("fiat rate %s-%s %d saved!\n", r.TickerSource, r.TickerTarget, r.Rate)
	return
}

func (p *DbmsProvider) SelectRangeCryptoRates(start, end, limit, offset int16) (rs *[]Rate, total int16, err error) {
	log.Println("select crypto rate by range", start, end, limit, offset)

	err = p.Db.Select(&rs,
		`select * from crypto where time >= ? and time <= ? offset ? limit ?`,
		start, end, offset, limit,
	)
	if err != nil {
		log.Println("select crypto rate by range err:", err)
		return
	}

	log.Println("crypto rate by range selected!")
	return
}

func (p *DbmsProvider) SelectRangeFiatRates(start, end string, limit, offset int16) (rs *[]Rate, total int16, err error) {
	log.Println("select crypto rate by range", start, end, limit, offset)

	err = p.Db.Select(&rs,
		`select * from crypto where time >= ? and time <= ? offset ? limit ?`,
		start, end, offset, limit,
	)
	if err != nil {
		log.Println("select crypto rate by range err:", err)
		return
	}

	log.Println("crypto rate by range selected!")
	return
}
