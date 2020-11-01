package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello can handle http requests
type Hello struct {
	l *log.Logger
}

// NewHello creates a new instance of the Hello struct with a specific logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("received a request")

	d, err := ioutil.ReadAll((r.Body))

	if err != nil {
		http.Error(rw, "Oops!", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello, %s!\n", d)
}
