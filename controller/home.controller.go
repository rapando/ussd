package controller

import (
	"fmt"
	"net/http"
)

func (b *Base) HomeController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for key, value := range r.Form {
		fmt.Printf("%s = %s\n", key, value)
	}
	response := "END This is the end"
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	w.Write([]byte(response))
}
