package main

import (
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})); err != nil {
		log.Fatal(err)
	}
}
