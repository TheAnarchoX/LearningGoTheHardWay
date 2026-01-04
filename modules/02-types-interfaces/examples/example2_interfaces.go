// Package examples demonstrates Go's interface system for experienced developers.
//
// This file shows:
// - Interface definition and implicit satisfaction
// - Multiple types implementing the same interface
// - Interface composition
// - Empty interface (interface{} / any)
// - Type assertions and type switches
package examples

import (
	"fmt"
	"math"
)

// Shape is an interface that defines behavior for geometric shapes.
//
// Coming from Java: No "implements" keyword needed - implicit satisfaction.
// Coming from Python: Like Protocol or ABC, but checked at compile time.
// Coming from C++: Like abstract base classes but without inheritance.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle represents a circle shape.
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle.
// By having this method, Circle automatically satisfies the Shape interface.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the perimeter of a circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle represents a rectangle shape.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of a rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Triangle represents a triangle shape.
type Triangle struct {
	Base   float64
	Height float64
	SideA  float64
	SideB  float64
	SideC  float64
}

// Area calculates the area of a triangle.
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter calculates the perimeter of a triangle.
func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

// PrintShapeInfo demonstrates interface usage.
// Any type that implements Shape can be passed here.
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// InterfaceBasics demonstrates basic interface usage.
func InterfaceBasics() {
	fmt.Println("=== Interface Basics ===")

	// Create different shapes
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}
	triangle := Triangle{Base: 10, Height: 8, SideA: 6, SideB: 8, SideC: 10}

	fmt.Println("Circle:")
	PrintShapeInfo(circle)

	fmt.Println("Rectangle:")
	PrintShapeInfo(rectangle)

	fmt.Println("Triangle:")
	PrintShapeInfo(triangle)

	// Can store different types in a slice of interfaces
	shapes := []Shape{circle, rectangle, triangle}
	fmt.Println("\nAll shapes:")
	for i, shape := range shapes {
		fmt.Printf("Shape %d: ", i+1)
		PrintShapeInfo(shape)
	}
}

// Stringer is similar to Java's toString() or Python's __str__.
// Many Go types implement fmt.Stringer for custom string representation.
type Stringer interface {
	String() string
}

// String implements the Stringer interface for Circle.
func (c Circle) String() string {
	return fmt.Sprintf("Circle(radius=%.2f)", c.Radius)
}

// String implements the Stringer interface for Rectangle.
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(width=%.2f, height=%.2f)", r.Width, r.Height)
}

// CommonInterfaces demonstrates using standard library interfaces.
func CommonInterfaces() {
	fmt.Println("\n=== Common Interfaces ===")

	circle := Circle{Radius: 5}
	// fmt.Println automatically uses String() method if available
	fmt.Println(circle)

	rectangle := Rectangle{Width: 10, Height: 5}
	fmt.Println(rectangle)
}

// ShapeReader and ShapeWriter demonstrate interface composition.
//
// Coming from Java: Like extending multiple interfaces.
// Coming from Go: Interfaces can embed other interfaces.
type ShapeReader interface {
	Read(p []byte) (n int, err error)
}

type ShapeWriter interface {
	Write(p []byte) (n int, err error)
}

// ShapeReadWriter combines ShapeReader and ShapeWriter interfaces.
type ShapeReadWriter interface {
	ShapeReader
	ShapeWriter
}

// InterfaceComposition demonstrates embedding interfaces.
func InterfaceComposition() {
	fmt.Println("\n=== Interface Composition ===")
	fmt.Println("ShapeReadWriter interface embeds both ShapeReader and ShapeWriter")
	fmt.Println("Any type implementing both Read and Write satisfies ShapeReadWriter")
}

// Describer is an interface for things that can describe themselves.
type Describer interface {
	Describe() string
}

// Describe implements Describer for Circle.
func (c Circle) Describe() string {
	return fmt.Sprintf("A circle with radius %.2f", c.Radius)
}

// TypeAssertion demonstrates type assertions.
//
// Coming from Java: Like instanceof and casting.
// Coming from Python: Like isinstance() and type checking.
func TypeAssertion() {
	fmt.Println("\n=== Type Assertion ===")

	var s Shape = Circle{Radius: 5}

	// Type assertion - check if Shape is actually a Circle
	if circle, ok := s.(Circle); ok {
		fmt.Printf("It's a circle with radius: %.2f\n", circle.Radius)
	} else {
		fmt.Println("Not a circle")
	}

	// Unsafe type assertion (panics if wrong type)
	// circle := s.(Circle)  // Use with caution!

	// Check for specific interface
	var shape Shape = Circle{Radius: 5}
	if describer, ok := shape.(Describer); ok {
		fmt.Printf("Description: %s\n", describer.Describe())
	}
}

// TypeSwitch demonstrates type switches for handling multiple types.
//
// Coming from Java: Cleaner than multiple instanceof checks.
// Coming from Python: Like match-case with type patterns (Python 3.10+).
func TypeSwitch() {
	fmt.Println("\n=== Type Switch ===")

	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 10, Height: 5},
		Triangle{Base: 10, Height: 8, SideA: 6, SideB: 8, SideC: 10},
	}

	for i, shape := range shapes {
		fmt.Printf("Shape %d: ", i+1)

		// Type switch
		switch v := shape.(type) {
		case Circle:
			fmt.Printf("Circle with radius %.2f\n", v.Radius)
		case Rectangle:
			fmt.Printf("Rectangle with dimensions %.2f x %.2f\n", v.Width, v.Height)
		case Triangle:
			fmt.Printf("Triangle with base %.2f and height %.2f\n", v.Base, v.Height)
		default:
			fmt.Printf("Unknown shape type\n")
		}
	}
}

// EmptyInterface demonstrates the empty interface (interface{} or any).
//
// Coming from Java: Like Object.
// Coming from TypeScript: Like 'any'.
// Coming from Python: Like typing.Any.
func EmptyInterface() {
	fmt.Println("\n=== Empty Interface ===")

	// interface{} can hold any value
	var anything interface{}

	anything = 42
	fmt.Printf("Int: %v (type: %T)\n", anything, anything)

	anything = "hello"
	fmt.Printf("String: %v (type: %T)\n", anything, anything)

	anything = Circle{Radius: 5}
	fmt.Printf("Circle: %v (type: %T)\n", anything, anything)

	anything = []int{1, 2, 3}
	fmt.Printf("Slice: %v (type: %T)\n", anything, anything)

	// Go 1.18+ introduced 'any' as an alias for interface{}
	var something any = "new syntax"
	fmt.Printf("Any: %v (type: %T)\n", something, something)
}

// ProcessValue demonstrates working with empty interface values.
func ProcessValue(v interface{}) {
	switch value := v.(type) {
	case int:
		fmt.Printf("Integer: %d, doubled: %d\n", value, value*2)
	case string:
		fmt.Printf("String: %q, length: %d\n", value, len(value))
	case Circle:
		fmt.Printf("Circle: area: %.2f\n", value.Area())
	case []int:
		sum := 0
		for _, n := range value {
			sum += n
		}
		fmt.Printf("Slice: %v, sum: %d\n", value, sum)
	default:
		fmt.Printf("Unknown type: %T\n", value)
	}
}

// EmptyInterfaceWithTypeSwitch demonstrates type switches with empty interface.
func EmptyInterfaceWithTypeSwitch() {
	fmt.Println("\n=== Empty Interface with Type Switch ===")

	values := []interface{}{
		42,
		"hello",
		Circle{Radius: 5},
		[]int{1, 2, 3, 4, 5},
		true,
	}

	for _, v := range values {
		ProcessValue(v)
	}
}

// NilInterface demonstrates nil interface gotcha.
//
// Important: An interface is nil only if both type and value are nil.
func NilInterface() {
	fmt.Println("\n=== Nil Interface ===")

	// Truly nil interface
	var s1 Shape
	fmt.Printf("s1 == nil: %v (s1: %v)\n", s1 == nil, s1)

	// Non-nil interface with nil value - GOTCHA!
	var c *Circle
	var s2 Shape = c
	fmt.Printf("s2 == nil: %v (s2: %v)\n", s2 == nil, s2)
	fmt.Printf("s2 is *Circle with nil value - calling methods will panic!\n")

	// Correct nil check for interface
	if s2 != nil {
		// This branch executes because s2 is not nil (it has type *Circle)
		// But calling methods will panic because the value is nil
		fmt.Printf("s2 has type %T\n", s2)
	}
}

// Polymorphism demonstrates polymorphic behavior with interfaces.
func Polymorphism() {
	fmt.Println("\n=== Polymorphism ===")

	// Different types, same interface
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 10, Height: 5},
	}

	totalArea := 0.0
	for _, shape := range shapes {
		// Each type implements Area() differently
		totalArea += shape.Area()
	}

	fmt.Printf("Total area: %.2f\n", totalArea)
}
