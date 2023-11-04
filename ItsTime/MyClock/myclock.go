package myclock

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondsHand(tm time.Time) Point {
	// scale, translate
	p := SecondsHandPositionInUnitCircle(tm)
	p.X = 90 * p.X
	p.Y = -90 * p.Y
	p.X += 150
	p.Y += 150
	return p
}

func SecondsHandPositionInUnitCircle(tm time.Time) Point {
	angle := SecondsInRadians(tm)
	return Point{math.Sin(angle), math.Cos(angle)}
}

func SecondsInRadians(tm time.Time) float64 {
	return math.Pi / (30 / float64(tm.Second()))
}
