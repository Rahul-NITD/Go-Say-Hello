package main

import (
	"fmt"
	"os"

	poker "github.com/Rahul-NITD/Poker"
)

func main() {
	fmt.Println("Let's Play Poker")
	fmt.Println("Type '{Name} wins' to record a win")
	store := poker.NewInMemoryStorage()
	alerter := poker.BlindAlerterFunc(poker.StdOutAlerter)
	game := poker.NewTexasHoldem(alerter, &store)
	cli := poker.NewCLI(os.Stdin, nil, game)
	cli.PlayPoker()
}
