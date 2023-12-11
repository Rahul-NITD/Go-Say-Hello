package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const secondHandLength = 90
const clockCenterX = 150
const clockCenterY = 150

func SecondsHandPoint(tm time.Time) Point {
	angle := ConvertSecondsToRadians(tm)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func SecondsHand(w io.Writer, tm time.Time) Point {
	p := SecondsHandPoint(tm)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
	return p
}

func ConvertSecondsToRadians(tm time.Time) float64 {
	return math.Pi / (30 / float64(tm.Second()))
}

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	SecondsHand(w, t)
	io.WriteString(w, svgEnd)
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
