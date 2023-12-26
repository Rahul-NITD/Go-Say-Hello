package poker

import (
	"log"
	"time"
)

type Game struct {
	alerter BlindAlerter
	store   PokerStorage
}

func NewGame(alerter BlindAlerter, store PokerStorage) Game {
	return Game{
		alerter: alerter,
		store:   store,
	}
}

func (g *Game) Start(numberOfPlayers int) {
	if g.alerter != nil {
		blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

		blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
		blindTime := 0 * time.Second
		for _, blind := range blinds {
			g.alerter.ScheduleAlertAfter(blindTime, blind)
			blindTime = blindTime + blindIncrement
		}
	} else {
		log.Println("No alerter passed")
	}
}

func (g *Game) Finish(winner string) {
	g.store.RecordWin(winner)
}
