package generics_test

import "testing"

// we want to be able to DRY our code concerned with equality
// we don't want to compare Apples to Oranges

func TestAssertFunctions(t *testing.T) {
	t.Run("Test Integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
}

func AssertEqual(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("got %d != %d", got, want)
	}
}

func AssertNotEqual(t testing.TB, got, want int) {
	if got == want {
		t.Errorf("got %d == %d", got, want)
	}
}
