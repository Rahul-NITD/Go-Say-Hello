package easyeasydi

import (
	"bytes"
	"testing"
)

func TestPrint(t *testing.T) {
	t.Run("Test Printing", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Rahul")
		got := buffer.String()
		want := "Hello, Rahul!\n"
		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	})
}
