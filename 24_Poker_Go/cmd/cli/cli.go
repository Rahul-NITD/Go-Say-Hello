package cli

import (
	"io"

	poker "github.com/Rahul-NITD/Poker"
)

type CLI struct {
	Store poker.PokerStorage
	Inp   io.Reader
}

func (c *CLI) PlayPoker() {
	c.Store.RecordWin("Chris")
}
