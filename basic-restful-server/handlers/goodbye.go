package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(log *log.Logger) *GoodBye {
	return &GoodBye{log}
}

func (h *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("GoodBye, World!")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Body: %s\n", d)

}
