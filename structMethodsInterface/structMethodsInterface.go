package structmethodsinterface

import "math"

func RectanglePerimeter(object Rectangle) float64 {
	return 2 * (object.Height + object.Width)
}

func RectangleArea(object Rectangle) float64 {
	return object.Height * object.Width
}

func CircleArea(object Circle) float64 {
	return (math.Pi * (object.Radius * object.Radius))
}
