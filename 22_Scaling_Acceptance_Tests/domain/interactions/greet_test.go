package interactions_test

import (
	"testing"

	"github.com/Rahul-NITD/scalingacceptancetests/domain/interactions"
	"github.com/Rahul-NITD/scalingacceptancetests/specs"
)

func TestGreet(t *testing.T) {
	t.Run("Test Greet", func(t *testing.T) {
		specs.GreeterSpecification(t, specs.GreetAdapter(interactions.Greet))
	})
	t.Run("Test No Name", func(t *testing.T) {
		got := interactions.Greet("")
		want := "Hello World"
		if got != want {
			t.Errorf("got %s != %s", got, want)
		}
	})
}
