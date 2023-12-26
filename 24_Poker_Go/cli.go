package poker

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
)

type CLI struct {
	store   PokerStorage
	inp     *bufio.Scanner
	alerter BlindAlerter
	out     io.Writer
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
	alertBlinds(c, players)

	c.inp.Scan()
	c.store.RecordWin(strings.Replace(c.inp.Text(), " wins", "", 1))
}

func alertBlinds(c *CLI, numberOfPlayers int) {
	if c.alerter != nil {
		blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

		blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
		blindTime := 0 * time.Second
		for _, blind := range blinds {
			c.alerter.ScheduleAlertAfter(blindTime, blind)
			blindTime = blindTime + blindIncrement
		}
	} else {
		log.Println("No alerter passed")
	}
}

func NewCLI(store PokerStorage, inp io.Reader, alerter BlindAlerter, out io.Writer) *CLI {
	return &CLI{
		store:   store,
		inp:     bufio.NewScanner(inp),
		alerter: alerter,
		out:     out,
	}
}
