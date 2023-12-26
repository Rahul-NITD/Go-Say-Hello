package poker

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type CLI struct {
	inp  *bufio.Scanner
	out  io.Writer
	game Game
}

func (c *CLI) PlayPoker() {

	if c.out != nil {
		fmt.Fprint(c.out, "Enter Number of Players : ")
	} else {
		log.Println("Output channel is nil")
	}

	c.inp.Scan()
	txt := strings.TrimSuffix(c.inp.Text(), "\n")
	players, err := strconv.Atoi(txt)
	if err != nil {
		log.Fatalf("Could not convert to integer. Run again %q", txt)
	}
	c.game.Start(players)

	c.inp.Scan()
	c.game.Finish(strings.Replace(c.inp.Text(), " wins", "", 1))
}

func NewCLI(store PokerStorage, inp io.Reader, alerter BlindAlerter, out io.Writer) *CLI {
	return &CLI{
		inp: bufio.NewScanner(inp),
		out: out,
		game: Game{
			alerter: alerter,
			store:   store,
		},
	}
}
