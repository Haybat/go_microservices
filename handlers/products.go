package handlers

import (
	"haybat.org/go_microservices/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	productList := data.GetProducts()
	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode json", http.StatusInternalServerError)
	}
}

/*
func (p *Products) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	productList := data.GetProducts()
	data, err := json.Marshal(productList)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	rw.Write(data)
}
*/
