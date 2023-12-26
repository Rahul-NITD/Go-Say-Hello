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
		inp := strings.NewReader("7\nChris wins\n")
		out := &bytes.Buffer{}
		store := NewSTUBStorage()
		alerter := SpyAlerter{}
		game := poker.NewTexasHoldem(&alerter, &store)
		cli := poker.NewCLI(inp, out, game)
		cli.PlayPoker()

		got := out.String()
		want := poker.NumberOfPlayersText

		if got != want {
			t.Error("Did not ask for number of players")
		}

	})

	t.Run("Test Game does not start when invalid input given", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Error("Start shouldn't be called")
		}

		got := out.String()
		want := poker.NumberOfPlayersText + poker.CannotConvertText

		if got != want {
			t.Errorf("got %q != %q", got, want)
		}

	})

	t.Run("Test Game stops when invalid input given", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := strings.NewReader("5\nThis game is good!")
		game := &GameSpy{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		got := out.String()
		want := poker.NumberOfPlayersText + poker.IncorrectInputText

		if got != want {
			t.Errorf("got %q != %q", got, want)
		}

	})

}

func assertPlayerWin(t testing.TB, store poker.PokerStorage, winner string, wantedWins int) {
	t.Helper()
	got, err := store.GetScore(winner)
	if err != nil {
		t.Fatalf("Error occured : %v", err)
	}
	want := wantedWins
	if got != want {
		t.Errorf("Did not record correct wins, got %d != %d", got, want)
	}
}
