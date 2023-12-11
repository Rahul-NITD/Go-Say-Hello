package itstime_test

import (
	myclock "GoSayHello/ItsTime/MyClock"
	"bytes"
	"encoding/xml"
	"testing"
	"time"
)

func TestClockSecondsHand(t *testing.T) {

	cases := []struct {
		name string
		time time.Time
		want myclock.Line
	}{
		{
			"Test at Midnight",
			time.Date(337, time.January, 1, 0, 0, 0, 0, time.UTC),
			myclock.Line{X1: 150, Y1: 150, X2: 150, Y2: 60},
		},
		{
			"Test at 00:00:30",
			time.Date(337, time.January, 1, 0, 0, 30, 0, time.UTC),
			myclock.Line{X1: 150, Y1: 150, X2: 150, Y2: 240},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			tm := test.time
			b := bytes.Buffer{}
			myclock.BuildSVG(&b, tm)
			svg := myclock.ClockSVG{}
			xml.Unmarshal(b.Bytes(), &svg)
			want := test.want
			if !ContainsLine(want, svg.Line) {
				t.Errorf("Required line %+v does not exist", want)
			}
		})
	}
}

func TestClockMinutesHand(t *testing.T) {

	cases := []struct {
		name string
		time time.Time
		want myclock.Line
	}{
		{
			"Test at Midnight",
			time.Date(337, time.January, 1, 0, 0, 0, 0, time.UTC),
			myclock.Line{X1: 150, Y1: 150, X2: 150, Y2: 60},
		},
		// {
		// 	"Test at 00:00:30",
		// 	time.Date(337, time.January, 1, 0, 0, 30, 0, time.UTC),
		// 	myclock.Line{X1: 150, Y1: 150, X2: 150, Y2: 240},
		// },
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			tm := test.time
			b := bytes.Buffer{}
			myclock.BuildSVG(&b, tm)
			svg := myclock.ClockSVG{}
			xml.Unmarshal(b.Bytes(), &svg)
			want := test.want
			if !ContainsLine(want, svg.Line) {
				t.Errorf("Required line %+v does not exist", want)
			}
		})
	}
}

func ContainsLine(want myclock.Line, got []myclock.Line) bool {
	for _, line := range got {
		if line == want {
			return true
		}
	}
	return false
}
