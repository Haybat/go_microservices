package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const hostUrl string = "localhost:9090"

func main() {

	http.HandleFunc("/hello", func(responseWriter http.ResponseWriter, request *http.Request) {
		log.Println("Hello microservice.")

		data, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, "Error in service", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(responseWriter, "Hello %s", data)
	})

	http.HandleFunc("/goodbye", func(responseWriter http.ResponseWriter, request *http.Request) {
		log.Println("Goodbye go microservice.")
	})

	http.ListenAndServe(hostUrl, nil)
}
