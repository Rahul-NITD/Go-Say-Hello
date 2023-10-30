package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	t.Run("Test Count Down", func(t *testing.T) {
		buffer := bytes.Buffer{}
		ss := &SpySleeper{}
		CountDown(&buffer, ss)
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
		if ss.Calls != 3 {
			t.Errorf("Called %d times wanted %d times", ss.Calls, 4)
		}
	})

	t.Run("Test sleep before every count", func(t *testing.T) {
		sco := &SpyCountdownOperations{}
		CountDown(sco, sco)
		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write"}
		if !reflect.DeepEqual(want, sco.calls) {
			t.Errorf("Order is not same!")
		}
	})

}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
