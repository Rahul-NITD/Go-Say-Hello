package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
