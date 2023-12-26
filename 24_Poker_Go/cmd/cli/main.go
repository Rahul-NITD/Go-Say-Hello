package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/Rahul-NITD/Poker"
)

func main() {
	fmt.Println("Let's Play Poker")
	fmt.Println("Type '{Name} wins' to record a win")
	store, close, err := poker.NewDBStorage(false)
	defer close(store)
	if err != nil {
		log.Fatalf("DB Error, %v", err)
	}
	alerter := poker.BlindAlerterFunc(poker.Alerter)
	game := poker.NewTexasHoldem(alerter, store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
