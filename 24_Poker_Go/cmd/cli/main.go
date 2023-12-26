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
	cli := poker.NewCLI(&store, os.Stdin, nil)
	cli.PlayPoker()
}