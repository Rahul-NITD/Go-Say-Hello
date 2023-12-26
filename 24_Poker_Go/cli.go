package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	Store PokerStorage
	Inp   io.Reader
}

func (c *CLI) PlayPoker() {
	reader := bufio.NewScanner(c.Inp)
	reader.Scan()
	c.Store.RecordWin(strings.Replace(reader.Text(), " wins", "", 1))
}

func NewCLI(store PokerStorage, inp io.Reader) *CLI {
	return &CLI{
		Store: store,
		Inp:   inp,
	}
}
