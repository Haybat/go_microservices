package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	g.l.Println("Goodbye microservice")

	data, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(rw, "Error in service", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Goodbye %s", data)
}
