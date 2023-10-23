package loopies

import "testing"

func TestLoopy(t *testing.T) {
	t.Run("Chk function exists", func(t *testing.T) {
		repeated := Loopy("", 0)
		expected := ""
		assertHelper(t, repeated, expected)
	})

	t.Run("repeat a 4 times", func(t *testing.T) {
		got := Loopy("a", 4)
		want := "aaaa"
		assertHelper(t, got, want)
	})

	t.Run("repeat 'a' 5 times", func(t *testing.T) {
		got := Loopy("a", 5)
		want := "aaaaa"
		assertHelper(t, got, want)
	})
}

func assertHelper(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
