package helloworld_test

import (
	"testing"

	helloworld "github.com/ChatGPTTeachesMeLearn/HelloWorld"
)

func TestHello(t *testing.T) {
	t.Run("first test", func(t *testing.T) {
		got := helloworld.Greet("Rahul")
		want := "Hi Rahul"
		if got != want {
			t.Errorf("got %s != %s", got, want)
		}
	})
}
