package dbmsprovider

import (
	"log"
)

func (p *DbmsProvider) InsertCryptoRate(r *Rate) (err error) {
	log.Printf("save crypto rate %s-%s %g...\n", r.Base, r.Currency, r.Rate)
	res, err := p.Db.Exec(
		"insert into crypto (time, base, currency, rate) values ($1, $2, $3, $4) on conflict do nothing",
		r.Time,
		r.Base,
		r.Currency,
		r.Rate,
	)
	if err != nil {
		log.Printf("save fiat crypto %s-%s %g err: %s\n", r.Base, r.Currency, r.Rate, err)
		return
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		log.Printf("crypto rate %s-%s %g already exist\n", r.Base, r.Currency, r.Rate)
		return
	}

	p.LastCrypto = r
	log.Printf("crypto rate %s-%s %g saved!\n", r.Base, r.Currency, r.Rate)
	return
}

func (p *DbmsProvider) InsertFiatRate(r *Rate) (err error) {
	log.Printf("save fiat rate %s-%s %g...\n", r.Base, r.Currency, r.Rate)
	res, err := p.Db.Exec(
		"insert into fiat (time, base, currency, rate) values ($1, $2, $3, $4) on conflict do nothing",
		r.Time,
		r.Base,
		r.Currency,
		r.Rate,
	)
	if err != nil {
		log.Printf("save fiat rate %s-%s %g err: %s\n", r.Base, r.Currency, r.Rate, err)
		return
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		log.Printf("fiat rate %s-%s %g already exist\n", r.Base, r.Currency, r.Rate)
		return
	}

	log.Printf("fiat rate %s-%s %g saved!\n", r.Base, r.Currency, r.Rate)
	return
}

func (p *DbmsProvider) SelectRangeCryptoRates(start, end, offset, limit string) (*[]Rate, uint16, error) {
	log.Printf("select crypto rate by range start:%s end:%s offset:%s limit:%s ", start, end, offset, limit)

	rows, err := p.Db.Query(`select time, base, currency, rate, count(*) over() as total from crypto where time between to_timestamp($1) and to_timestamp($2) offset $3 limit $4`,
		start, end, offset, limit)
	if err != nil {
		log.Println("select crypto rate by range err:", err)
		return nil, 0, err
	}

	var rates []Rate
	var total uint16
	for rows.Next() {
		var rate Rate
		err = rows.Scan(&rate.Time, &rate.Base, &rate.Currency, &rate.Rate, &total)
		if err != nil {
			log.Println("select crypto rows.Scan err:", err)
			return nil, 0, err
		}
		rates = append(rates, rate)
	}
	err = rows.Err()
	if err != nil {
		log.Println("select crypto rows.Err:", err)
		return nil, 0, err
	}

	log.Println("crypto rate by range selected!")
	return &rates, total, nil
}

func (p *DbmsProvider) SelectRangeFiatRates(start, end string) (*[]Rate, error) {
	log.Printf("select crypto rate by range start:%s end:%s ", start, end)

	var rates []Rate
	err := p.Db.Select(&rates, `select * from fiat where time between $1 and $2`, start, end)
	if err != nil {
		log.Println("select fiat rate by range err:", err)
		return nil, err
	}

	log.Println("fiat rate by range selected!")
	return &rates, nil
}
