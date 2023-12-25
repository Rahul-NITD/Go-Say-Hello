package interactions_test

import (
	"testing"

	"github.com/Rahul-NITD/scalingacceptancetests/domain/interactions"
	"github.com/Rahul-NITD/scalingacceptancetests/specs"
)

func TestCurse(t *testing.T) {
	t.Run("Test Curse", func(t *testing.T) {
		specs.CurseSpecification(t, specs.CurseAdapter(interactions.Curse))
	})
	t.Run("Test Curse No Name", func(t *testing.T) {
		got := interactions.Curse("")
		want := "Go to Hell Nobody"
		if got != want {
			t.Errorf("got %s != %s", got, want)
		}
	})
}
