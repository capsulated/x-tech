package config

type Config struct {
	Api    Api
	Server Server
	Dbms   Dbms
	Cron   Cron
}

type Api struct {
	CryptoSourceUrl string `mapstructure:"API_CRYPTO_SOURCE_URL"`
	FiatSourceUrl   string `mapstructure:"API_FIAT_SOURCE_URL"`
}

type Server struct {
	Host string `mapstructure:"SERVER_HOST"`
	Port string `mapstructure:"SERVER_PORT"`
}

type Dbms struct {
	Driver   string `mapstructure:"DBMS_DRIVER"`
	Host     string `mapstructure:"DBMS_HOST"`
	Port     string `mapstructure:"DBMS_PORT"`
	Ssl      bool   `mapstructure:"DBMS_SSL"`
	User     string `mapstructure:"DBMS_USER"`
	Password string `mapstructure:"DBMS_PASSWORD"`
	DbName   string `mapstructure:"DBMS_DB_NAME"`
}

type Cron struct {
	Crypto string `mapstructure:"CRON_CRYPTO"`
	Fiat   string `mapstructure:"CRON_FIAT"`
}
