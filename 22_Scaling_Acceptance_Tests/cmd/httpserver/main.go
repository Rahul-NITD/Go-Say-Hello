package main

import (
	"log"
	"net/http"

	"github.com/Rahul-NITD/scalingacceptancetests/adapters/httpserver"
)

func main() {
	handler := httpserver.NewHandler()
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
