package main

import (
	"log"
	"net/http"
)

func startServer() {

	port := ":80"
	address := "127.0.0.1"

	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)
	log.Printf("Served at http://%s", address+port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
