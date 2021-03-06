package controller

import (
	"log"
	"net/http"
	"os"
	"github.com/rs/cors"

)

func (b *Base) Run() {
	port := ":" + os.Getenv("PORT")
	log.Printf("API listening on port %s", port)

	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(b.Router)
	log.Fatalln(http.ListenAndServe(port, handler))

}
