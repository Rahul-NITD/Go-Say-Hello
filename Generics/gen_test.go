package generics_test

import "testing"

// we want to be able to DRY our code concerned with equality
// we don't want to compare Apples to Oranges

func TestAssertFunctions(t *testing.T) {
	t.Run("Test Integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("Test strings", func(t *testing.T) {
		AssertEqual(t, "HI", "HI")
		AssertNotEqual(t, "HI", "BYE")
	})

	t.Run("Completely absurd", func(t *testing.T) {
		AssertEqual(t, "0", 0) // we dont want this to run
	})

}

func AssertEqual(t testing.TB, got, want interface{}) {
	if got != want {
		t.Errorf("got %d != %d", got, want)
	}
}

func AssertNotEqual(t testing.TB, got, want interface{}) {
	if got == want {
		t.Errorf("got %d == %d", got, want)
	}
}
