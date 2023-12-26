package poker_test

import (
	"strings"
	"testing"

	poker "github.com/Rahul-NITD/Poker"
)

func TestCLI(t *testing.T) {
	t.Run("Test for Chris", func(t *testing.T) {
		inp := strings.NewReader("Chris wins\n")
		store := NewSTUBStorage()
		cli := &poker.CLI{Store: &store, Inp: inp}
		cli.PlayPoker()
		assertPlayerWin(t, &store, "Chris", 1)
	})
	t.Run("Test for Adam", func(t *testing.T) {
		inp := strings.NewReader("Adam wins\n")
		store := NewSTUBStorage()
		cli := &poker.CLI{Store: &store, Inp: inp}
		cli.PlayPoker()
		assertPlayerWin(t, &store, "Adam", 1)
	})
}

func assertPlayerWin(t testing.TB, store poker.PokerStorage, winner string, wantedWins int) {
	t.Helper()
	if sc, err := store.GetScore(winner); sc != 1 || err != nil {
		t.Errorf("%s score : %d\nerror : %v", winner, sc, err)
	}
	got, err := store.GetScore(winner)
	if err != nil {
		t.Fatalf("Error occured : %v", err)
	}
	want := wantedWins
	if got != want {
		t.Errorf("Did not record correct wins, got %d != %d", got, want)
	}
}
