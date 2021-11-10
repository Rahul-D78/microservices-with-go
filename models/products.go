package models

import (
	"encoding/json"
	"io"
)

//we are using struct fields to modify the response
type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Products is a collection of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	{
		ID:   1,
		Name: "Legion",
	},
	{
		ID:   2,
		Name: "Razer",
	},
}
