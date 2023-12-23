package shapesout

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

func Area(r Rectangle) float64 {
	return r.width * r.height
}

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func (circ Circle) area() float64 {
	return math.Pi * circ.radius * circ.radius
}

func (trig Triangle) area() float64 {
	return 0.5 * trig.base * trig.height
}

type Shape interface {
	area() float64
}
