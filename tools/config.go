package tool

import (
	"github.com/globocom/config"
	"log"
	"os"
)

type Conf struct {
	Servername   string
	Serverport   string
	Cookiesecret string
	Dbaddress    string
	Dbname       string
}

func (conf *Conf)  Config(file string) error {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		err := Copy("config/registry.default.yml", "config/registry.yml")
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	if err := config.ReadConfigFile(file); err != nil {
		log.Fatal(err)
		return err
	}

	if value, err := config.Get("servername"); err != nil {
		conf.Servername = "localhost"
	} else {
		conf.Servername = value.(string)
	}
	if value, err := config.Get("serverport"); err != nil {
		conf.Serverport = "8989"
	} else {
		conf.Serverport = value.(string)
	}
	if value, err := config.Get("cookiesecret"); err != nil {
		conf.Cookiesecret = "uif23fui4iiuy3i5g4y23u"
	} else {
		conf.Cookiesecret = value.(string)
	}
	if value, err := config.Get("dbaddress"); err != nil {
		conf.Dbaddress = "localhost:28015"
	} else {
		conf.Dbaddress = value.(string)
	}
	if value, err := config.Get("dbname"); err != nil {
		conf.Dbname = "registry"
	} else {
		conf.Dbname = value.(string)
	}

	return nil
}
