package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	h.l.Println("Hello microservice.")

	data, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(rw, "Error in service", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", data)
}
