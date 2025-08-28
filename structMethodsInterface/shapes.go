package structmethodsinterface

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (R Rectangle) Area() float64 {
	return (R.Height * R.Width)
}

type Circle struct {
	radius float64
}

func (C Circle) Area() float64 {
	return (math.Pi * (C.radius * C.radius))
}
