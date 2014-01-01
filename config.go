package main

import (
	"github.com/globocom/config"
	"log"
	"os"
)

type Conf struct {
	servername string
	serverport string
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
		conf.serverport = ":4004"
	} else {
		conf.serverport = value.(string)
	}

	return nil
}
