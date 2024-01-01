package numbergame_test

// Overview?

import (
	numbergame "GoSayHello/02_Integers"
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("test 2 + 2", func(t *testing.T) {
		sum := numbergame.Adder(2, 2)
		expected := 4
		if sum != expected {
			t.Errorf("got '%d' but required '%d'", expected, sum)
		}
	})
	t.Run("test 4+7", func(t *testing.T) {
		sum := numbergame.Adder(4, 7)
		expected := 11
		if sum != expected {
			t.Errorf("got '%d' but required '%d'", expected, sum)
		}
	})
}

func ExampleAdder() {
	sum := numbergame.Adder(3, 4)
	fmt.Println(sum)
	// Output: 7
}
