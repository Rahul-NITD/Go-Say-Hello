package interactions_test

import (
	"testing"

	"github.com/Rahul-NITD/scalingacceptancetests/domain/interactions"
	"github.com/Rahul-NITD/scalingacceptancetests/specs"
)

func TestGreet(t *testing.T) {
	specs.GreeterSpecification(t, specs.GreetAdapter(interactions.Greet))
}
