package main

import (
	"context"
	"haybat.org/go_microservices/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, os.Kill)

	sig := <-signalChan
	l.Println("Received terminate, shutting down! ", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
