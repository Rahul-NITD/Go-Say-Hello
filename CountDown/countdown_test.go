package countdown

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	t.Run("Test Count Down", func(t *testing.T) {
		buffer := bytes.Buffer{}
		CountDown(buffer)
		got := buffer.String()
		want := "3"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
