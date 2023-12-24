package clockface_test

import (
	"GoSayHello/16_Maths/clockface"
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"
)

func TestClockFace(t *testing.T) {
	t.Run("Test seconds at 12", func(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
		got := clockface.SecondsHand(&bytes.Buffer{}, tm)
		want := clockface.Point{X: 150, Y: 150 - 90}
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Test seconds hand at 30 seconds", func(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)
		got := clockface.SecondsHand(&bytes.Buffer{}, tm)
		want := clockface.Point{X: 150, Y: 150 + 90}
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{SimpleTime(0, 0, 30), clockface.Point{0, -1}},
		{SimpleTime(0, 0, 0), clockface.Point{0, 1}},
		{SimpleTime(0, 0, 15), clockface.Point{1, 0}},
		{SimpleTime(0, 0, 45), clockface.Point{-1, 0}},
	}
	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := clockface.SecondsHandPoint(test.time)
			want := test.point
			if !roughlyEqualPoint(got, want) {
				t.Errorf("got %v wanted %v", got, want)
			}
		})
	}
}

func roughlyEqual(got, want float64) bool {
	const threshold = 1e-7
	return math.Abs(got-want) < threshold
}

func roughlyEqualPoint(got, want clockface.Point) bool {
	return roughlyEqual(got.X, want.X) && roughlyEqual(got.Y, want.Y)
}

func TestSecondsToRadianConversion(t *testing.T) {

	cases := []struct {
		time  time.Time
		angle float64
	}{
		{SimpleTime(0, 0, 0), 0},
		{SimpleTime(0, 0, 30), math.Pi},
		{SimpleTime(0, 0, 45), 2 * math.Pi / 4 * 3},
		{SimpleTime(0, 0, 45), 2 * math.Pi / 4 * 3},
		{SimpleTime(0, 0, 7), 2 * math.Pi / 60 * 7},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := clockface.ConvertSecondsToRadians(test.time)
			want := test.angle
			if got != want {
				t.Errorf("got %f want %f", got, want)
			}
		})
	}
}

func SimpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.January, 1, hours, minutes, seconds, 0, time.UTC)
}

func testName(tm time.Time) string {
	return tm.Format("15:04:05")
}

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSvgWriter(t *testing.T) {
	tm := SimpleTime(0, 0, 0)
	b := bytes.Buffer{}
	clockface.SVGWriter(&b, tm)
	svg := SVG{}
	xml.Unmarshal(b.Bytes(), &svg)

	want := Line{150, 150, 150, 0}

	for _, line := range svg.Line {
		if line == want {
			return
		}
	}
	t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", want, svg.Line)
}
