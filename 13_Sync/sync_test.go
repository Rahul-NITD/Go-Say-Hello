package amiinsync

import (
	"sync"
	"testing"
)

func assertCounterVal(t testing.TB, counter *Counter, expected int) {
	t.Helper()
	if counter.Value() != expected {
		t.Errorf("Got %d wanted %d", counter.val, expected)
	}
}

func TestCounter(t *testing.T) {
	t.Run("Test Counter Up Only", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounterVal(t, &counter, 3)
	})

	t.Run("Run it Concurrently", func(t *testing.T) {
		counter := Counter{}
		expected := 1000
		var wg sync.WaitGroup
		wg.Add(expected)
		for i := 0; i < expected; i++ {
			go func() {
				wg.Done()
				counter.Inc()
			}()
		}
		wg.Wait()
		assertCounterVal(t, &counter, expected)
	})
}
