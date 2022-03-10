package handlers

import (
	"basic-web-server/data"
	"encoding/json"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		p.getProducts(response, request)
	}
	response.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(response http.ResponseWriter, request *http.Request) {
	lp := data.GetProducts(
	data, err := json.Marshal(lp)
	if err != nil {
		http.Error(response, "Unable to marshal json", http.StatusInternalServerError)
	}
	response.Write(data)
}
