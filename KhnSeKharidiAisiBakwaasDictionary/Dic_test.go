package khnsekharidiaisibakwaasdictionary

import "testing"

func TestDic(t *testing.T) {

	assertHelper := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got '%s' want '%s'", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("Did not expect an error here")
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("Got '%s' want '%s'", got.Error(), want.Error())
		}
		if got == nil {
			t.Fatal("Expected an error")
		}
	}

	t.Run("Test Dic", func(t *testing.T) {
		dictionary := Dictionary{"Hello": "It is a greeting"}
		got, err := dictionary.Search("Hello")
		want := dictionary["Hello"]
		assertNoError(t, err)
		assertHelper(t, got, want)
	})

	t.Run("Unknown Error", func(t *testing.T) {
		dic := Dictionary{"Hello": "Greeting"}
		_, err := dic.Search("Test")
		assertError(t, err, ErrSearchWordNotFound)
	})

}
