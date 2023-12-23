package numbergame

import "testing"

func TestAdder(t *testing.T) {
	t.Run("test 2 + 2", func(t *testing.T) {
		sum := Adder(2, 2)
		expected := 4
		if sum != expected {
			t.Errorf("got '%d' but required '%d'", expected, sum)
		}
	})
	t.Run("test 4+7", func(t *testing.T) {
		sum := Adder(4, 7)
		expected := 11
		if sum != expected {
			t.Errorf("got '%d' but required '%d'", expected, sum)
		}
	})
}
