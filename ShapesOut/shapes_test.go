package shapesout

import "testing"

func TestShapes(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, hasArea float64) {
		got := shape.area()
		if got != hasArea {
			t.Errorf("%#v got %.2f want %.2f", shape, got, hasArea)
		}
	}

	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{12.0, 4.0}, hasArea: 48.0},
		{shape: Circle{10.0}, hasArea: 314.1592653589793},
		{shape: Triangle{12.0, 6.0}, hasArea: 36.0},
	}

	for _, testCase := range areaTests {
		checkArea(t, testCase.shape, testCase.hasArea)
	}

}
