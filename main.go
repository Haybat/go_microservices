package main

import (
	"haybat.org/go_microservices/handlers"
	"log"
	"net/http"
	"os"
)

const hostUrl string = "localhost:9090"

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/hello", helloHandler)
	sm.Handle("/goodbye", goodbyeHandler)

	http.ListenAndServe(hostUrl, sm)
}
