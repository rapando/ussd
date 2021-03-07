package controller

import (
	"log"
	"net/http"
)

func (b *Base) HomeController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for key, value := range r.Form {
		log.Printf("%s = %s\n", key, value)
	}
	response := "END This is the end"
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("connection", "close")
	w.WriteHeader(200)
	w.Write([]byte(response))
}
