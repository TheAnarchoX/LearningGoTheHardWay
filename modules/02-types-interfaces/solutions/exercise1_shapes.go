// Package solutions contains correct implementations of the exercises.
//
// SOLUTION 1: Shapes with Interfaces
//
// This file shows the correct implementation with all bugs fixed.
//
// Key fixes:
// - Circle.Area() uses correct formula: π * r²
// - Circle.Name() method implemented
// - Square.Perimeter() uses correct formula: 4 * side
// - Triangle.Perimeter() calculates hypotenuse correctly
// - Ellipse uses value receivers for consistency
// - FindLargestShape() compares areas correctly
// - ScaleShape() handles all shape types
// - CalculateStats() initializes MinArea correctly
// - CompareShapes() has correct comparison logic
package solutions

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
// FIXED: Correct formula π * r²
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the perimeter of a circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Name returns the name of the shape.
// FIXED: Implemented missing method
func (c Circle) Name() string {
	return "Circle"
}

// Square represents a square.
type Square struct {
	Side float64
}

// Area calculates the area of a square.
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// Perimeter calculates the perimeter of a square.
// FIXED: Correct formula 4 * side
func (s Square) Perimeter() float64 {
	return 4 * s.Side
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
// FIXED: Calculates hypotenuse using Pythagorean theorem
func (t Triangle) Perimeter() float64 {
	hypotenuse := math.Sqrt(t.Base*t.Base + t.Height*t.Height)
	return t.Base + t.Height + hypotenuse
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
// FIXED: Using value receiver for consistency
func (e Ellipse) Area() float64 {
	return math.Pi * e.MajorAxis * e.MinorAxis
}

// Perimeter approximates the perimeter of an ellipse using Ramanujan's formula.
// FIXED: Using value receiver for consistency
func (e Ellipse) Perimeter() float64 {
	h := math.Pow(e.MajorAxis-e.MinorAxis, 2) / math.Pow(e.MajorAxis+e.MinorAxis, 2)
	return math.Pi * (e.MajorAxis + e.MinorAxis) * (1 + 3*h/(10+math.Sqrt(4-3*h)))
}

// Name returns the name of the shape.
// FIXED: Using value receiver for consistency
func (e Ellipse) Name() string {
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
// FIXED: Properly handles empty slice and compares areas correctly
func FindLargestShape(shapes []Shape) Shape {
	if len(shapes) == 0 {
		return nil
	}

	largest := shapes[0]
	largestArea := largest.Area()

	// FIXED: Compare areas, not indices
	for i := 1; i < len(shapes); i++ {
		area := shapes[i].Area()
		if area > largestArea {
			largest = shapes[i]
			largestArea = area
		}
	}
	return largest
}

// ScaleShape creates a new shape that is scaled by the given factor.
// FIXED: Handles all shape types
func ScaleShape(s Shape, factor float64) Shape {
	switch shape := s.(type) {
	case Circle:
		return Circle{Radius: shape.Radius * factor}
	case Square:
		return Square{Side: shape.Side * factor}
	case Triangle:
		// FIXED: Added case for Triangle
		return Triangle{
			Base:   shape.Base * factor,
			Height: shape.Height * factor,
		}
	case Ellipse:
		return Ellipse{
			MajorAxis: shape.MajorAxis * factor,
			MinorAxis: shape.MinorAxis * factor,
		}
	default:
		// For unknown shapes, return the original
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
// FIXED: Properly initializes MinArea
func CalculateStats(shapes []Shape) ShapeStats {
	if len(shapes) == 0 {
		return ShapeStats{}
	}

	// FIXED: Initialize MinArea with first shape's area or MaxFloat64
	stats := ShapeStats{
		Count:   len(shapes),
		MinArea: math.MaxFloat64,
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
// FIXED: Correct comparison logic
func CompareShapes(s1, s2 Shape) int {
	area1 := s1.Area()
	area2 := s2.Area()

	// FIXED: Correct return values
	if area1 < area2 {
		return -1
	} else if area1 > area2 {
		return 1
	}
	return 0
}
