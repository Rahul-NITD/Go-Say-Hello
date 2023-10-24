package khnsekharidiaisibakwaasdictionary

import "testing"

func TestDic(t *testing.T) {

	assertStrings := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("search a word", func(t *testing.T) {
		dictionary := map[string]string{"test": "this is test"}

		got := Search(dictionary, "test")
		want := "this is test"

		assertStrings(t, got, want)

	})
}
