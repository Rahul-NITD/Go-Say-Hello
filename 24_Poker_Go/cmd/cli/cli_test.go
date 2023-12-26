package cli_test

import (
	"strings"
	"testing"

	poker "github.com/Rahul-NITD/Poker"
	"github.com/Rahul-NITD/Poker/cmd/cli"
)

func TestCLI(t *testing.T) {
	inp := strings.NewReader("Chris wins\n")
	store := poker.NewInMemoryStorage()
	cli := &cli.CLI{Store: &store, Inp: inp}
	cli.PlayPoker()
	assertPlayerWin(t, &store, "Chris", 1)
}

func assertPlayerWin(t testing.TB, store poker.PokerStorage, winner string, wantedWins int) {
	t.Helper()
	if sc, err := store.GetScore("Chris"); sc != 1 || err != nil {
		t.Errorf("%s score : %d\nerror : %v", winner, sc, err)
	}
	got, err := store.GetScore("Chris")
	if err != nil {
		t.Fatalf("Error occured : %v", err)
	}
	want := wantedWins
	if got != want {
		t.Errorf("Did not record correct wins, got %d != %d", got, want)
	}
}
