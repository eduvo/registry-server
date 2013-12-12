package main

import (
	"github.com/globocom/config"
	"log"
	"os"
)

type Conf struct {
	servername string
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
	value, err := config.Get("servername")
	str := value.(string)
	conf.servername = str

	return err
}
