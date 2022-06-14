package config

type Config struct {
	Api    Api
	Server Server
	Dbms   Dbms
	Cron   Cron
}

type Api struct {
	CryptoSourceUrl string
	FiatSourceUrl   string
}

type Server struct {
	Host string
	Port string
}

type Dbms struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type Cron struct {
	Crypto string
	Fiat   string
}
