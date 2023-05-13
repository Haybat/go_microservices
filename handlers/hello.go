package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Hello is a simple handler
type Hello struct {
	l *log.Logger
}

// NewHello Creates a new hello handler with the given logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHTTP implements the go http.Handler interface
func (h *Hello) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	h.l.Println("Handle Hello request.")

	data, err := io.ReadAll(request.Body)
	if err != nil {
		h.l.Println("Error reading body", err)
		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", data)
}
