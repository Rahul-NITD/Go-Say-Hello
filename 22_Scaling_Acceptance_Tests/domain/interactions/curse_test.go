package interactions_test

import (
	"testing"

	"github.com/Rahul-NITD/scalingacceptancetests/domain/interactions"
	"github.com/Rahul-NITD/scalingacceptancetests/specs"
)

func TestCurse(t *testing.T) {
	specs.CurseSpecification(t, specs.CurseAdapter(interactions.Curse))
}
