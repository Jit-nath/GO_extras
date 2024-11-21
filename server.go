package main

import (
	"log"
	"net/http"
)

func startServer() {

	var port string = ":8080"
	var address string = "127.0.0.1"

	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)
	log.Printf("Served at http://%s", address+port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
