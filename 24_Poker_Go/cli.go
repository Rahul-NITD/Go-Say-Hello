package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type CLI struct {
	store   PokerStorage
	inp     io.Reader
	alerter BlindAlerter
}

func (c *CLI) PlayPoker() {
	if c.alerter != nil {
		blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
		blindTime := 0 * time.Second
		for _, blind := range blinds {
			c.alerter.ScheduleAlertAfter(blindTime, blind)
			blindTime = blindTime + 10*time.Minute
		}
	}
	reader := bufio.NewScanner(c.inp)
	reader.Scan()
	c.store.RecordWin(strings.Replace(reader.Text(), " wins", "", 1))
}

func NewCLI(store PokerStorage, inp io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{
		store:   store,
		inp:     inp,
		alerter: alerter,
	}
}
