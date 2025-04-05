package server

import (
	"log"
	"net/http"
)

func StartServer(metricsHandler http.Handler, address string) {
	http.Handle("/metrics", metricsHandler)
	log.Printf("Starting server on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
