package main

import (
	"log"
	"net/http"

	"github.com/Rahul-NITD/SAGo/PokerServer/pokerhttpserver"
)

func main() {
	server := pokerhttpserver.PlayerServer{
		Storage: pokerhttpserver.NewInMemoryStorage(),
	}

	log.Fatal(http.ListenAndServe(":8000", http.Handler(&server)))
}
