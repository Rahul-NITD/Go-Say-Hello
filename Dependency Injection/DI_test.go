package easyeasydi

import (
	"bytes"
	"testing"
)

func TestXxx(t *testing.T) {
	t.Run("Test Printing", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Rahul")
		got := buffer.String()
		want := "Hello, Rahul!"
		if got != want {
			t.Errorf("got %s wanted %s", got, want)
		}
	})
}
