package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

func NewConfig() (c Config) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	viper.SetConfigFile(filepath.Dir(wd) + "/.env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	api := &Api{}
	err = viper.Unmarshal(&api)
	if err != nil {
		log.Fatalln(err)
	}

	srv := &Server{}
	err = viper.Unmarshal(&srv)
	if err != nil {
		log.Fatalln(err)
	}

	db := &Dbms{}
	err = viper.Unmarshal(&db)
	if err != nil {
		log.Fatalln(err)
	}

	crn := &Cron{}
	err = viper.Unmarshal(&crn)
	if err != nil {
		log.Fatalln(err)
	}

	return Config{*api, *srv, *db, *crn}
}
