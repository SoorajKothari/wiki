package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "wiki/pkg/config"
)

var globalConfig *Config

func Start(config *Config) {
	globalConfig = config
	router := mux.NewRouter()
	router.HandleFunc("v1/book", GetBooks).Methods("GET")
	router.HandleFunc("v1/book", AddBook).Methods("POST")

	err := http.ListenAndServe(":"+config.ServerPort, router)
	if err != nil {
		log.Fatal(err)
	}
}
