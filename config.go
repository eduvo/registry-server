package main

import (
	"github.com/globocom/config"
	"log"
	"os"
)

type Conf struct {
	servername   string
	serverport   string
	cookiesecret string
	dbaddress    string
	dbname       string
}

func Config() error {
	if _, err := os.Stat(ConfigFile); os.IsNotExist(err) {
		err := Copy("registry.default.yml", "registry.yml")
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	if err := config.ReadConfigFile(ConfigFile); err != nil {
		log.Fatal(err)
		return err
	}

	if value, err := config.Get("servername"); err != nil {
		conf.servername = "localhost"
	} else {
		conf.servername = value.(string)
	}
	if value, err := config.Get("serverport"); err != nil {
		conf.serverport = "8989"
	} else {
		conf.serverport = value.(string)
	}
	if value, err := config.Get("cookiesecret"); err != nil {
		conf.cookiesecret = "uif23fui4iiuy3i5g4y23u"
	} else {
		conf.cookiesecret = value.(string)
	}
	if value, err := config.Get("dbaddress"); err != nil {
		conf.dbaddress = "localhost:28015"
	} else {
		conf.dbaddress = value.(string)
	}
	if value, err := config.Get("dbname"); err != nil {
		conf.dbname = "registry"
	} else {
		conf.dbname = value.(string)
	}

	return nil
}
