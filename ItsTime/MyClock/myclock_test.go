package myclock_test

import (
	myclock "GoSayHello/ItsTime/MyClock"
	"math"
	"testing"
	"time"
)

func assertPoint(t testing.TB, got, want myclock.Point) {
	t.Helper()
	if !(NearlyEquals(got.X, want.X) && NearlyEquals(got.Y, want.Y)) {
		t.Errorf("got %+v want %+v", got, want)
	}
}

func NearlyEquals(a, b float64) bool {
	return math.Abs(a-b) < 1e-7
}

func assertFloat(t testing.TB, got, want float64) {
	t.Helper()
	if !NearlyEquals(got, want) {
		t.Errorf("got %.3f want %.3f", got, want)
	}
}

func SimpleTime(hh, mm, ss int) time.Time {
	return time.Date(1337, time.January, 1, hh, mm, ss, 0, time.UTC)
}

func SmolTestName(t time.Time) string {
	return t.Format("15:04:05")
}

// func TestSecondsHand(t *testing.T) {
// 	t.Run("Integration Test: Check Seconds Hand at 00:00:00", func(t *testing.T) {
// 		tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

// 	})
// }
