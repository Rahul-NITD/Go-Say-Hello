package cli_test

import (
	"testing"

	poker "github.com/Rahul-NITD/Poker"
	"github.com/Rahul-NITD/Poker/cmd/cli"
)

func TestCLI(t *testing.T) {
	store := poker.NewInMemoryStorage()
	cli := &cli.CLI{Store: &store}
	cli.PlayPoker()
	if sc, err := store.GetScore("Adam"); sc != 1 || err != nil {
		t.Errorf("Adam score : %d\nerror : %v", sc, err)
	}
}
