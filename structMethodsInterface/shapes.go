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
	Radius float64
}

func (C Circle) Area() float64 {
	return (math.Pi * (C.Radius * C.Radius))
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	Height float64
	Base   float64
}

func (t Triangle) Area() float64 {
	return (0.5 * t.Base * t.Height)
}

// Here are the key reasons why interfaces are so useful:

// 1. Polymorphism and Abstraction
// Interfaces allow you to treat different types as a single, common type, as long as they implement the same methods. This is known as polymorphism. This lets you write functions that can operate on a variety of types, making your code more abstract and less coupled to specific implementations.

// For example, imagine a function CalculateTotalArea(shapes []Shape) that can sum the areas of different shapes like rectangles, circles, and triangles. Without interfaces, you would need to write a separate function for each shape type, or use more complex and less "Go-idiomatic" approaches.

// 2. Decoupling and Loose Coupling
// Interfaces help decouple your code. A component that uses an interface doesn't need to know the concrete type it's interacting with; it only needs to know the methods it can call. This separation of concerns makes your code easier to manage, test, and modify. You can swap out one implementation for another without affecting the rest of the program, as long as the new implementation satisfies the interface. This is especially useful for testing, where you can replace a complex, external dependency (like a database connection) with a simple mock implementation that satisfies the same interface.

// 3. Dependency Injection
// Interfaces are the foundation of dependency injection in Go. Instead of a function creating its own dependencies, you can "inject" them via the function's parameters. By using interfaces for these dependencies, you make the function more flexible and testable.

// For example, a function that needs to write to a file could take an io.Writer interface as a parameter. This allows you to pass in a *os.File, a bytes.Buffer, or even a mock for testing, without changing the function's code.
