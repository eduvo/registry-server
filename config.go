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

	return nil
}
