package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the struct for an API product
// struct tags!! `field tags` changing the name of the output
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"` // omit from output
	DeletedOn   string  `json:"-"`
}

func(p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p) // decoder returns error
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w) // instead of using Marshal: tihs is faster for not too big files
	return e.Encode(p)      // encode returns an error
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID() // from the productList var
	productList = append(productList, p)
}

func getNextID() int {
	lp := productList[len(productList) - 1] // get last product
	return lp.ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Pr 1",
		Description: "product 1 description",
		Price:       2.45,
		SKU:         "ABC123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Pr 2",
		Description: "product 2 description",
		Price:       22.22,
		SKU:         "ABC2222",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
