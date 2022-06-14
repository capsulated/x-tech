package config

import "os"

func NewConfig() (c Config) {
	api := &Api{
		os.Getenv("API_CRYPTO_SOURCE_URL"),
		os.Getenv("API_FIAT_SOURCE_URL"),
	}

	srv := &Server{
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	}

	db := &Dbms{
		os.Getenv("DBMS_DRIVER"),
		os.Getenv("DBMS_HOST"),
		os.Getenv("DBMS_PORT"),
		os.Getenv("DBMS_USER"),
		os.Getenv("DBMS_PASSWORD"),
		os.Getenv("DBMS_DB_NAME"),
	}

	crn := &Cron{
		os.Getenv("CRON_CRYPTO"),
		os.Getenv("CRON_FIAT"),
	}

	return Config{*api, *srv, *db, *crn}
}
