// Package exercises contains practice problems with intentional bugs.
//
// EXERCISE 1: Shapes with Interfaces
//
// Fix the bugs in this file to make all tests pass.
// The bugs are marked with // BUG: comments.
//
// Learning objectives:
// - Implement interfaces correctly
// - Understand method receivers (value vs pointer)
// - Work with interface types
// - Fix type satisfaction issues
package exercises

import (
	"fmt"
	"math"
)

// Shape is an interface that all geometric shapes must implement.
type Shape interface {
	Area() float64
	Perimeter() float64
	Name() string
}

// Circle represents a circle.
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle.
// BUG: Formula is incorrect - should be π * r²
func (c Circle) Area() float64 {
	return math.Pi * c.Radius
}

// Perimeter calculates the perimeter of a circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Name returns the name of the shape.
// BUG: Missing implementation - Circle doesn't implement Shape interface
// TODO: Implement this method to return "Circle"

// Square represents a square.
type Square struct {
	Side float64
}

// Area calculates the area of a square.
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// Perimeter calculates the perimeter of a square.
// BUG: Wrong calculation - should be 4 * side
func (s Square) Perimeter() float64 {
	return 2 * s.Side
}

// Name returns the name of the shape.
func (s Square) Name() string {
	return "Square"
}

// Triangle represents a right triangle.
type Triangle struct {
	Base   float64
	Height float64
}

// Area calculates the area of a triangle.
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter calculates the perimeter of a triangle.
// BUG: Wrong calculation - for a right triangle, need to calculate hypotenuse
// The formula should use Pythagorean theorem: √(base² + height²)
func (t Triangle) Perimeter() float64 {
	return t.Base + t.Height
}

// Name returns the name of the shape.
func (t Triangle) Name() string {
	return "Triangle"
}

// Ellipse represents an ellipse.
type Ellipse struct {
	MajorAxis float64
	MinorAxis float64
}

// Area calculates the area of an ellipse.
// BUG: Using pointer receiver when other shapes use value receivers
// This causes issues with interface satisfaction consistency
func (e *Ellipse) Area() float64 {
	return math.Pi * e.MajorAxis * e.MinorAxis
}

// Perimeter approximates the perimeter of an ellipse using Ramanujan's formula.
// BUG: Using pointer receiver - should be consistent with Area
func (e *Ellipse) Perimeter() float64 {
	h := math.Pow(e.MajorAxis-e.MinorAxis, 2) / math.Pow(e.MajorAxis+e.MinorAxis, 2)
	return math.Pi * (e.MajorAxis + e.MinorAxis) * (1 + 3*h/(10+math.Sqrt(4-3*h)))
}

// Name returns the name of the shape.
// BUG: Using pointer receiver - should be consistent
func (e *Ellipse) Name() string {
	return "Ellipse"
}

// PrintShapeInfo prints information about a shape.
func PrintShapeInfo(s Shape) {
	fmt.Printf("%s: Area = %.2f, Perimeter = %.2f\n", s.Name(), s.Area(), s.Perimeter())
}

// TotalArea calculates the total area of all shapes.
func TotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// FindLargestShape returns the shape with the largest area.
// BUG: Returns nil if shapes slice is empty, but should handle this case
// TODO: Return an error or a zero-value shape instead of nil
func FindLargestShape(shapes []Shape) Shape {
	if len(shapes) == 0 {
		return nil
	}

	largest := shapes[0]
	// BUG: Wrong comparison - should compare areas, not indices
	for i := 1; i < len(shapes); i++ {
		if i > 0 {
			largest = shapes[i]
		}
	}
	return largest
}

// ScaleShape creates a new shape that is scaled by the given factor.
// BUG: Type switch doesn't handle all shape types
func ScaleShape(s Shape, factor float64) Shape {
	switch shape := s.(type) {
	case Circle:
		return Circle{Radius: shape.Radius * factor}
	case Square:
		return Square{Side: shape.Side * factor}
	// BUG: Missing case for Triangle
	// TODO: Add case for Triangle
	default:
		// BUG: Returns original shape instead of scaled version
		return s
	}
}

// ShapeStats holds statistics about a collection of shapes.
type ShapeStats struct {
	Count     int
	TotalArea float64
	MinArea   float64
	MaxArea   float64
	AvgArea   float64
}

// CalculateStats calculates statistics for a slice of shapes.
// BUG: Doesn't initialize MinArea properly - it will be 0.0 instead of the first shape's area
func CalculateStats(shapes []Shape) ShapeStats {
	if len(shapes) == 0 {
		return ShapeStats{}
	}

	stats := ShapeStats{
		Count:   len(shapes),
		MinArea: 0.0, // BUG: Should be math.MaxFloat64 or first shape's area
		MaxArea: 0.0,
	}

	for _, shape := range shapes {
		area := shape.Area()
		stats.TotalArea += area

		if area < stats.MinArea {
			stats.MinArea = area
		}
		if area > stats.MaxArea {
			stats.MaxArea = area
		}
	}

	// BUG: Integer division instead of float division
	stats.AvgArea = stats.TotalArea / float64(len(shapes))

	return stats
}

// CompareShapes compares two shapes by area.
// Returns:
//
//	-1 if s1 has smaller area than s2
//	 0 if s1 has equal area to s2
//	 1 if s1 has larger area than s2
//
// BUG: Comparison logic is reversed
func CompareShapes(s1, s2 Shape) int {
	area1 := s1.Area()
	area2 := s2.Area()

	if area1 < area2 {
		return 1 // BUG: Should return -1
	} else if area1 > area2 {
		return -1 // BUG: Should return 1
	}
	return 0
}
