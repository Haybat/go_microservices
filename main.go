package main

import (
	"haybat.org/go_microservices/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

const hostUrl string = "localhost:9090"

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/hello", helloHandler)
	sm.Handle("/goodbye", goodbyeHandler)

	server := &http.Server{
		Addr:         hostUrl,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	server.ListenAndServe()
}
