package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/sideroff/go-microservices/data"
)

// Products handler
type Products struct {
	l *log.Logger
}

// NewProducts constructor
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// attach function serveHTTP to any p: Products
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Request:", r.Method, r.URL.Path)

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	} else if r.Method == http.MethodPost {
		p.createProduct(rw, r)
		return
	} else if r.Method == http.MethodPut {

		//		   eg. /products/3
		requestPath := r.URL.Path

		rgx := regexp.MustCompile("/([0-9]+)")

		groups := rgx.FindAllStringSubmatch(requestPath, -1)

		if len(groups) != 1 || len(groups[0]) != 2 {
			http.Error(rw, "Please provide a valid id.", http.StatusBadRequest)
			return
		}

		idString := groups[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw, "Bad URL.", http.StatusBadRequest)
		}

		p.l.Println("got id", id)
		p.updateProduct(id, rw, r)
		return
	}

	http.Error(rw, "Request method not allowed.", http.StatusMethodNotAllowed)

}

// attach function getProducts to any p: Products
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Could not encode data.", http.StatusInternalServerError)
	}
}

func (p *Products) createProduct(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Could not decode data. Please provide valid json.", http.StatusBadRequest)
	}

	p.l.Printf("Prod %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Could not decode data. Please provide valid json.", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)

	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found.", http.StatusNotFound)
	} else if err != nil {
		http.Error(rw, "Internal server error.", http.StatusInternalServerError)
	}
	return
}
