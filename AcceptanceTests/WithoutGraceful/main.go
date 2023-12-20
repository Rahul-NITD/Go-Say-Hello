package main

import (
	"log"
	"net/http"

	"github.com/quii/go-graceful-shutdown/acceptancetests"
)

func main() {
	httpServer := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(acceptancetests.SlowHandler)}

	// server := gracefulshutdown.NewServer(httpServer)

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("uh oh, didnt shutdown gracefully, some responses may have been lost %v", err)
	}

	log.Println("shutdown gracefully! all responses were sent")
}
