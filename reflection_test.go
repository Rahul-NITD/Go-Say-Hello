package reflection_test

import (
	"fmt"
	"testing"
	"testing/quick"
)

type Input struct {
	name string
	num  int
}

func TestReflection(t *testing.T) {
	t.Run("Test Property based", func(t *testing.T) {
		assertion := func(name string, num int) bool {
			x := Input{
				name: name,
				num:  num,
			}
			fmt.Printf("%+v\n", x)
			return true
		}
		if err := quick.Check(assertion, nil); err != nil {
			t.Error(err)
		}
	})
}
