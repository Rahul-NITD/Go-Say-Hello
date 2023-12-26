package cli

import poker "github.com/Rahul-NITD/Poker"

type CLI struct {
	Store poker.PokerStorage
}

func (c *CLI) PlayPoker() {
	c.Store.RecordWin("Adam")
}
