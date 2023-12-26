package poker

import (
	"io"
	"log"
	"time"
)

type TexasHoldem struct {
	alerter BlindAlerter
	store   PokerStorage
}

type Game interface {
	Start(numberOfPlayers int, alertDest io.Writer)
	Finish(winner string)
}

func NewTexasHoldem(alerter BlindAlerter, store PokerStorage) TexasHoldem {
	return TexasHoldem{
		alerter: alerter,
		store:   store,
	}
}

func (g TexasHoldem) Start(numberOfPlayers int, alertDest io.Writer) {
	if g.alerter != nil {
		blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

		blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
		blindTime := 0 * time.Second
		for _, blind := range blinds {
			g.alerter.ScheduleAlertAfter(blindTime, blind, alertDest)
			blindTime = blindTime + blindIncrement
		}
	} else {
		log.Println("No alerter passed")
	}
}

func (g TexasHoldem) Finish(winner string) {
	g.store.RecordWin(winner)
}
