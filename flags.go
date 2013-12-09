package main

import (
	"flag"
)

var (
	ConfigFile string
)

func init() {
	flag.StringVar(&ConfigFile, "config", "registry.yml", "Path to the Registry config file")
}

func Flags() error {
	if !flag.Parsed() {
		flag.Parse()
	}
	return nil
}
