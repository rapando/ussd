package controller

import "github.com/gorilla/mux"

func (b *Base) InitializeRouter() {
	b.Router = mux.NewRouter()

	b.Router.HandleFunc("/", b.HomeController).Methods("POST")
}