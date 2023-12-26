package poker_test

import (
	"bytes"
	"strings"
	"testing"

	poker "github.com/Rahul-NITD/Poker"
)

func TestCLI(t *testing.T) {
	t.Run("Test for Chris", func(t *testing.T) {
		inp := strings.NewReader("5\nChris wins\n")
		store := NewSTUBStorage()
		game := poker.NewTexasHoldem(nil, &store)
		cli := poker.NewCLI(inp, nil, game)
		cli.PlayPoker()
		assertPlayerWin(t, &store, "Chris", 1)
	})
	t.Run("Test for Adam", func(t *testing.T) {
		inp := strings.NewReader("5\nAdam wins\n")
		store := NewSTUBStorage()
		game := poker.NewTexasHoldem(nil, &store)
		cli := poker.NewCLI(inp, nil, game)
		cli.PlayPoker()
		assertPlayerWin(t, &store, "Adam", 1)
	})

	t.Run("Test it prompts for number of users", func(t *testing.T) {
		inp := strings.NewReader("7\n")
		out := &bytes.Buffer{}
		store := NewSTUBStorage()
		alerter := SpyAlerter{}
		game := poker.NewTexasHoldem(&alerter, &store)
		cli := poker.NewCLI(inp, out, game)
		cli.PlayPoker()

		got := out.String()
		want := "Enter Number of Players : "

		if got != want {
			t.Error("Did not ask for number of players")
		}

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
