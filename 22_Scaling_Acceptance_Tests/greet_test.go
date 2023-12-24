package scalingacceptancetests_test

import (
	"testing"

	"github.com/Rahul-NITD/scalingacceptancetests"
	"github.com/Rahul-NITD/scalingacceptancetests/specs"
)

func TestGreet(t *testing.T) {
	specs.GreeterSpecification(t, specs.GreetAdapter(scalingacceptancetests.Greet))
}
