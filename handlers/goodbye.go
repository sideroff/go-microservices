package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Goodbye can handle http requests
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye creates a new instance of the Hello struct with a specific logger
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("received a goodbye")

	d, err := ioutil.ReadAll((r.Body))

	if err != nil {
		http.Error(rw, "Oops!", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Goodbye, %s!\n", d)
}
