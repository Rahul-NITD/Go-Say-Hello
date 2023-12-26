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

const (
	NumberOfPlayersText = "Enter Number of Players : "
	CannotConvertText   = "Could not convert to integer. Run again"
	IncorrectInputText  = "Incorrect input."
)

func (c *CLI) PlayPoker() {

	if c.out != nil {
		fmt.Fprint(c.out, NumberOfPlayersText)
	} else {
		log.Println("Output channel is nil")
	}

	c.inp.Scan()
	txt := strings.TrimSuffix(c.inp.Text(), "\n")
	players, err := strconv.Atoi(txt)
	if err != nil {
		fmt.Fprint(c.out, CannotConvertText)
		return
	}
	c.game.Start(players, c.out)

	c.inp.Scan()
	if !strings.Contains(c.inp.Text(), "wins") {
		fmt.Fprint(c.out, IncorrectInputText)
	}
	c.game.Finish(strings.Replace(c.inp.Text(), " wins", "", 1))
}

func NewCLI(inp io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		inp:  bufio.NewScanner(inp),
		out:  out,
		game: game,
	}
}
