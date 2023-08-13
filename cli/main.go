package main

import (
	"log"
	. "wiki/pkg/config"
	"wiki/pkg/router"
)

var conf *Config

func init() {
	log.Println("Initializing...!")

	conf, err := GetConfig()
	if err != nil {
		log.Println("Error while Initializing", err)
		return
	}

	log.Println("Initialize with db Name: ", conf.DbName)
}

func main() {
	router.Start(conf)
}
