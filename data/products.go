package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:-`
	UpdatedOn   string  `json:-`
	DeletedOn   string  `json:-`
}

type Products []*Product

var ErrorProductNotFound = fmt.Errorf("Product not found")

func GetProducts() Products {
	return productList
}

func AddProduct(product *Product) {
	product.ID = getNextId()
	productList = append(productList, product)
}

func UpdateProduct(id int, product *Product) error {
	oldProduct, index, err := getProductById(id)

	if err != nil {
		return err
	}

	fmt.Println(oldProduct.Name)

	product.ID = id
	productList[index] = product

	for _, p := range productList {
		fmt.Println(p.Name)
	}

	return nil
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func getProductById(id int) (*Product, int, error) {

	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrorProductNotFound
}

func getNextId() int {
	p := productList[len(productList)-1]
	return p.ID + 1
}

var productList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy Milky coffee",
		Price:       2.25,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short & strong coffee",
		Price:       1.99,
		SKU:         "fdj34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
