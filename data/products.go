package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product struct defines the structure of an api product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products interface
type Products []*Product

// FromJSON Reads from r and decodes from json
func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

// ToJSON Encodes to json and writes to w
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latter",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "SKU1",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong coffee without milk",
		Price:       1.45,
		SKU:         "SKU2",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

// GetProducts returns products
func GetProducts() Products {
	return productList
}

// AddProduct adds a product to the list
func AddProduct(p *Product) {
	p.ID = getNextID()

	productList = append(productList, p)
}

// UpdateProduct udpates a product based on the received id
func UpdateProduct(id int, p *Product) error {

	_, index, err := findProduct(id)

	if err != nil {
		return err
	}

	productList[index] = p
	return nil
}

// ErrorProductNotFound error
var ErrorProductNotFound = fmt.Errorf("product not found")

func findProduct(id int) (*Product, int, error) {
	for i, product := range productList {
		if product.ID == id {
			return product, i, nil
		}
	}

	return nil, -1, ErrorProductNotFound
}

func getNextID() int {
	return productList[len(productList)-1].ID + 1
}
