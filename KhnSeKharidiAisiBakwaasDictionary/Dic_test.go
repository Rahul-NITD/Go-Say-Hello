package khnsekharidiaisibakwaasdictionary

import "testing"

func TestDic(t *testing.T) {

	assertHelper := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got '%s' want '%s'", got, want)
		}
	}

	t.Run("Test Dic", func(t *testing.T) {
		dictionary := Dictionary{"Hello": "It is a greeting"}
		got, err := dictionary.Search("Hello")
		want := dictionary["Hello"]
		if err != nil {
			t.Fatal("Did not expect an error here")
		}
		assertHelper(t, got, want)
	})

	t.Run("Unknown Error", func(t *testing.T) {
		dic := Dictionary{"Hello": "Greeting"}
		_, err := dic.Search("Test")
		want := "search word not found in Dictionary"
		if err == nil {
			t.Fatal("Expected an Error here")
		}
		assertHelper(t, err.Error(), want)
	})

}
