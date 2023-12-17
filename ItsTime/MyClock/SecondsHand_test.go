package myclock_test

import (
	myclock "GoSayHello/ItsTime/MyClock"
	"math"
	"testing"
	"time"
)

func TestSecondsHand(t *testing.T) {
	t.Run("UT1: Test Position in unit circle", func(t *testing.T) {
		cases := []struct {
			time  time.Time
			point myclock.Point
		}{
			{SimpleTime(0, 0, 00), myclock.Point{0, 1}},
			{SimpleTime(0, 0, 30), myclock.Point{0, -1}},
			{SimpleTime(0, 0, 45), myclock.Point{-1, 0}},
			{SimpleTime(0, 0, 15), myclock.Point{1, 0}},
		}
		for _, test := range cases {
			t.Run(SmolTestName(test.time), func(t *testing.T) {
				got := myclock.SecondsHandPositionInUnitCircle(test.time)
				want := test.point
				assertPoint(t, got, want)
			})
		}
	})
	t.Run("UT2: Test Seconds to Angles", func(t *testing.T) {
		cases := []struct {
			time  time.Time
			angle float64
		}{
			{SimpleTime(0, 0, 00), 0},
			{SimpleTime(0, 0, 30), math.Pi},
			{SimpleTime(0, 0, 15), math.Pi / 2},
			{SimpleTime(0, 0, 45), 3 * math.Pi / 2},
		}
		for _, test := range cases {
			t.Run(SmolTestName(test.time), func(t *testing.T) {
				got := myclock.SecondsInRadians(test.time)
				want := test.angle
				assertFloat(t, got, want)
			})
		}
	})
	t.Run("UT3: Test Seconds Hand Position", func(t *testing.T) {
		cases := []struct {
			time  time.Time
			point myclock.Point
		}{
			{SimpleTime(0, 0, 00), myclock.Point{150, 60}},
			{SimpleTime(0, 0, 30), myclock.Point{150, 150 + 90}},
			{SimpleTime(0, 0, 15), myclock.Point{150 + 90, 150}},
			{SimpleTime(0, 0, 45), myclock.Point{150 - 90, 150}},
		}
		for _, test := range cases {
			t.Run(SmolTestName(test.time), func(t *testing.T) {
				got := myclock.SecondsHand(nil, test.time)
				want := test.point
				assertPoint(t, got, want)
			})
		}
	})
}
