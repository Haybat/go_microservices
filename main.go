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

	sm := http.NewServeMux()
	sm.Handle("/hello", helloHandler)

	http.ListenAndServe(hostUrl, sm)
}
