package integers_test

import (
	"testing"
	"testing/quick"

	integers "github.com/ChatGPTTeachesMeLearn/Integers"
)

func TestSeriesSum(t *testing.T) {
	t.Run("Test until 5", func(t *testing.T) {
		got := integers.SumUpto(5)
		want := 15
		if got != want {
			t.Errorf("got %d != %d", got, want)
		}
	})
	t.Run("Test until 15", func(t *testing.T) {
		got := integers.SumUpto(15)
		want := 120
		if got != want {
			t.Errorf("got %d != %d", got, want)
		}
	})

	t.Run("Test 0", func(t *testing.T) {
		got := integers.SumUpto(0)
		want := 0
		if got != want {
			t.Errorf("got %d != %d", got, want)
		}
	})

	t.Run("Test negative", func(t *testing.T) {
		got := integers.SumUpto(-5)
		want := 0
		if got != want {
			t.Errorf("got %d != %d", got, want)
		}
	})

}

func PropertyTesting(t *testing.T) {
	t.Run("S + S = 2 * S", func(t *testing.T) {
		assertion := func(n int) bool {
			return integers.SumUpto(n)+integers.SumUpto(n) == 2*integers.SumUpto(n)
		}
		if err := quick.Check(assertion, nil); err != nil {
			t.Error("failed checks", err)
		}
	})
	t.Run("S(n) - S(n-1) = n", func(t *testing.T) {
		assertion := func(n int) bool {
			return integers.SumUpto(n)-integers.SumUpto(n-1) == n
		}
		if err := quick.Check(assertion, nil); err != nil {
			t.Error("failed checks", err)
		}
	})
}
