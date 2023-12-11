package myclock

import (
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

type ClockSVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

func BuildSVG(w io.Writer, tm time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	SecondsHand(w, tm)
	MinutesHand(w, tm)
	HoursHand(w, tm)
	io.WriteString(w, svgEnd)
}

func SecondsHand(w io.Writer, tm time.Time) Point {
	// scale, translate
	p := SecondsHandPositionInUnitCircle(tm)
	p.X = 90 * p.X
	p.Y = -90 * p.Y
	p.X += 150
	p.Y += 150
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
	return p
}

func SecondsHandPositionInUnitCircle(tm time.Time) Point {
	angle := SecondsInRadians(tm)
	return Point{math.Sin(angle), math.Cos(angle)}
}

func SecondsInRadians(tm time.Time) float64 {
	return math.Pi / (30 / float64(tm.Second()))
}

func MinutesHand(w io.Writer, tm time.Time) Point {
	// scale, translate
	p := MinutesHandPositionInUnitCircle(tm)
	p.X = 90 * p.X
	p.Y = -90 * p.Y
	p.X += 150
	p.Y += 150
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
	return p
}

func MinutesHandPositionInUnitCircle(tm time.Time) Point {
	angle := MinutesInRadians(tm)
	return Point{math.Sin(angle), math.Cos(angle)}
}

func MinutesInRadians(tm time.Time) float64 {
	return math.Pi/(30/float64(tm.Minute())) + SecondsInRadians(tm)/60
}

func HoursHand(w io.Writer, tm time.Time) Point {
	// scale, translate
	p := HoursHandPositionInUnitCircle(tm)
	p.X = 50 * p.X
	p.Y = -50 * p.Y
	p.X += 150
	p.Y += 150
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
	return p
}

func HoursHandPositionInUnitCircle(tm time.Time) Point {
	angle := HoursInRadians(tm)
	return Point{math.Sin(angle), math.Cos(angle)}
}

func HoursInRadians(tm time.Time) float64 {
	return math.Pi/(12/float64(tm.Hour())) + MinutesInRadians(tm)/60
}
