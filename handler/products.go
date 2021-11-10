package handler

import (
	"log"
	"net/http"

	"github.com/Rahul-D78/micro-go/models"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

//This is a http Handler so needs a ServeHTTP method
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}
	//catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	//GET request to products list using JSON encoding
	//READ more https://pkg.go.dev/encoding/json

	lp := models.GetProducts()
	err := lp.ToJSON(w)

	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}
