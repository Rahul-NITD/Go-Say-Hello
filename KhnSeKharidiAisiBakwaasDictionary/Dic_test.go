package khnsekharidiaisibakwaasdictionary

import "testing"

func TestDic(t *testing.T) {
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

	t.Run("Test add words", func(t *testing.T) {
		dic := Dictionary{"Hello": "Greeting"}
		key := "Test"
		value := "This is a test"
		addErr := dic.Add(key, value)
		assertNoError(t, addErr)
		got, err := dic.Search(key)
		assertNoError(t, err)
		assertHelper(t, got, value)
	})

	t.Run("Test add word exists", func(t *testing.T) {
		pvalue := "Greeting"
		key := "Hello"
		value := "Again??"
		dic := Dictionary{key: pvalue}
		addErr := dic.Add(key, value)
		assertError(t, addErr, ErrKeyAlreadyExist)
		got, err := dic.Search(key)
		assertNoError(t, err)
		assertHelper(t, got, pvalue)

	})

}

var assertHelper = func(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got '%s' want '%s'", got, want)
	}
}

var assertNoError = func(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Did not expect an error here")
	}
}

var assertError = func(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("Got '%q' want '%q'", got, want)
	}
	if got == nil {
		t.Fatal("Expected an error")
	}
}

func TestUpdate(t *testing.T) {
	t.Run("test update word", func(t *testing.T) {
		dic := Dictionary{"Hello": "Greeting"}
		upderr := dic.Update("Hello", "Updated Greeting")
		assertNoError(t, upderr)
		got, err := dic.Search("Hello")
		assertNoError(t, err)
		assertHelper(t, got, "Updated Greeting")
	})
	t.Run("test update word missing", func(t *testing.T) {
		dic := Dictionary{"Hello": "Greeting"}
		upderr := dic.Update("Test", "This is Test")
		assertError(t, upderr, ErrCannotUpdateNonExistingKey)
		_, err := dic.Search("Test")
		assertError(t, err, ErrSearchWordNotFound)
	})
}
