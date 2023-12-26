package main

import (
	"log"
	"net/http"

	poker "github.com/Rahul-NITD/Poker"
)

func main() {
	storage, close, _ := poker.NewDBStorage(false)
	defer close(storage)
	alerter := poker.BlindAlerterFunc(poker.Alerter)
	game := poker.NewTexasHoldem(alerter, storage)
	server := poker.NewServer(storage, game)
	log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(server.ServeHTTP)))
}
