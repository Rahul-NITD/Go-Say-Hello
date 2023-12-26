package cli

import (
	"bufio"
	"io"
	"strings"

	poker "github.com/Rahul-NITD/Poker"
)

type CLI struct {
	Store poker.PokerStorage
	Inp   io.Reader
}

func (c *CLI) PlayPoker() {
	reader := bufio.NewScanner(c.Inp)
	reader.Scan()
	c.Store.RecordWin(strings.Replace(reader.Text(), " wins", "", 1))
}
