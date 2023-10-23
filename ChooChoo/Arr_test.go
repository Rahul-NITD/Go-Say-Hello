package choochoo

import "testing"

func TestSum(t *testing.T) {
	t.Run("test for 5 elements", func(t *testing.T) {
		numbers := []int{1, 4, 2, 5, 6}
		got := Sum(numbers)
		want := 18
		assertHelper(t, got, want, numbers)

	})
}

func assertHelper(t testing.TB, got, want int, values []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d for '%v'", got, want, values)
	}
}
