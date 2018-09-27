package main

import (
	"log"
	"net/http"
	"os"

	h "./handlers"
)

var (
	defaultPort = "8080"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}
	http.HandleFunc("/put/event", h.PutEvent)
	http.HandleFunc("^/", h.Acknowledge)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
