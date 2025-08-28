package structmethodsinterface

import "testing"

func TestPerimeter(t *testing.T) {
	givenRectangle := Rectangle{10.0, 10.0}
	calculatedPerimeter := RectanglePerimeter(givenRectangle)
	expectedResults := 40.0

	if calculatedPerimeter != expectedResults {
		t.Errorf("the calculated permeter came out to be %.2f but the expected was %.2f", calculatedPerimeter, expectedResults)
	}
}

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		// got := RectangleArea(rectangle)
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		// got := CircleArea(circle)
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
