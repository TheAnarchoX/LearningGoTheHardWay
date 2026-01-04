package exercises

import (
	"math"
	"testing"
)

// TestCircleArea verifies Circle area calculation.
func TestCircleArea(t *testing.T) {
	circle := Circle{Radius: 5}
	area := circle.Area()
	expected := math.Pi * 25 // π * r²

	if math.Abs(area-expected) > 0.001 {
		t.Errorf("Circle area incorrect: expected %.2f, got %.2f", expected, area)
	}
}

// TestCircleImplementsShape verifies Circle implements Shape.
func TestCircleImplementsShape(t *testing.T) {
	circle := Circle{Radius: 5}
	var s Shape = circle

	if s == nil {
		t.Error("Circle should implement Shape interface")
	}

	// Should be able to call Name()
	name := s.Name()
	if name != "Circle" {
		t.Errorf("Expected name 'Circle', got %q", name)
	}
}

// TestSquarePerimeter verifies Square perimeter calculation.
func TestSquarePerimeter(t *testing.T) {
	square := Square{Side: 10}
	perimeter := square.Perimeter()
	expected := 40.0 // 4 * side

	if perimeter != expected {
		t.Errorf("Square perimeter incorrect: expected %.2f, got %.2f", expected, perimeter)
	}
}

// TestTrianglePerimeter verifies Triangle perimeter calculation.
func TestTrianglePerimeter(t *testing.T) {
	// Right triangle with sides 3, 4, 5
	triangle := Triangle{Base: 3, Height: 4}
	perimeter := triangle.Perimeter()
	hypotenuse := math.Sqrt(3*3 + 4*4) // = 5
	expected := 3 + 4 + hypotenuse     // = 12

	if math.Abs(perimeter-expected) > 0.001 {
		t.Errorf("Triangle perimeter incorrect: expected %.2f, got %.2f", expected, perimeter)
	}
}

// TestEllipseWithValueReceiver verifies Ellipse works with value receivers.
func TestEllipseWithValueReceiver(t *testing.T) {
	// This test will fail if Ellipse uses pointer receivers
	ellipse := Ellipse{MajorAxis: 10, MinorAxis: 5}

	// Should work with value type
	var s Shape = ellipse

	if s == nil {
		t.Error("Ellipse should implement Shape interface with value receiver")
	}

	area := s.Area()
	if area <= 0 {
		t.Errorf("Ellipse area should be positive, got %.2f", area)
	}
}

// TestPrintShapeInfo verifies PrintShapeInfo function.
func TestPrintShapeInfo(t *testing.T) {
	circle := Circle{Radius: 5}
	// Should not panic
	PrintShapeInfo(circle)
}

// TestTotalArea verifies total area calculation.
func TestTotalArea(t *testing.T) {
	shapes := []Shape{
		Circle{Radius: 5},
		Square{Side: 10},
		Triangle{Base: 10, Height: 8},
	}

	total := TotalArea(shapes)
	expectedCircle := math.Pi * 25
	expectedSquare := 100.0
	expectedTriangle := 40.0
	expected := expectedCircle + expectedSquare + expectedTriangle

	if math.Abs(total-expected) > 0.001 {
		t.Errorf("Total area incorrect: expected %.2f, got %.2f", expected, total)
	}
}

// TestFindLargestShape verifies finding the largest shape.
func TestFindLargestShape(t *testing.T) {
	shapes := []Shape{
		Circle{Radius: 3},            // Area ≈ 28.27
		Square{Side: 10},             // Area = 100
		Triangle{Base: 5, Height: 4}, // Area = 10
	}

	largest := FindLargestShape(shapes)
	if largest == nil {
		t.Fatal("FindLargestShape returned nil")
	}

	// Should be the square
	if largest.Name() != "Square" {
		t.Errorf("Expected largest shape to be Square, got %s", largest.Name())
	}
}

// TestFindLargestShapeEmpty verifies handling of empty slice.
func TestFindLargestShapeEmpty(t *testing.T) {
	shapes := []Shape{}
	largest := FindLargestShape(shapes)

	// Should handle empty slice gracefully (not panic)
	if largest != nil {
		t.Error("FindLargestShape should return nil for empty slice")
	}
}

// TestScaleShape verifies shape scaling.
func TestScaleShape(t *testing.T) {
	// Test Circle scaling
	circle := Circle{Radius: 5}
	scaled := ScaleShape(circle, 2.0)
	scaledCircle, ok := scaled.(Circle)
	if !ok {
		t.Fatal("Scaled shape should be Circle")
	}
	if scaledCircle.Radius != 10.0 {
		t.Errorf("Expected radius 10.0, got %.2f", scaledCircle.Radius)
	}

	// Test Square scaling
	square := Square{Side: 10}
	scaled = ScaleShape(square, 0.5)
	scaledSquare, ok := scaled.(Square)
	if !ok {
		t.Fatal("Scaled shape should be Square")
	}
	if scaledSquare.Side != 5.0 {
		t.Errorf("Expected side 5.0, got %.2f", scaledSquare.Side)
	}

	// Test Triangle scaling
	triangle := Triangle{Base: 10, Height: 8}
	scaled = ScaleShape(triangle, 2.0)
	scaledTriangle, ok := scaled.(Triangle)
	if !ok {
		t.Fatal("Scaled shape should be Triangle")
	}
	if scaledTriangle.Base != 20.0 || scaledTriangle.Height != 16.0 {
		t.Errorf("Expected base 20.0 and height 16.0, got base %.2f and height %.2f",
			scaledTriangle.Base, scaledTriangle.Height)
	}
}

// TestCalculateStats verifies statistics calculation.
func TestCalculateStats(t *testing.T) {
	shapes := []Shape{
		Circle{Radius: 5},             // Area ≈ 78.54
		Square{Side: 10},              // Area = 100
		Triangle{Base: 10, Height: 8}, // Area = 40
	}

	stats := CalculateStats(shapes)

	if stats.Count != 3 {
		t.Errorf("Expected count 3, got %d", stats.Count)
	}

	expectedTotal := math.Pi*25 + 100 + 40
	if math.Abs(stats.TotalArea-expectedTotal) > 0.001 {
		t.Errorf("Expected total area %.2f, got %.2f", expectedTotal, stats.TotalArea)
	}

	// MinArea should be 40 (Triangle)
	if math.Abs(stats.MinArea-40.0) > 0.001 {
		t.Errorf("Expected min area 40.0, got %.2f", stats.MinArea)
	}

	// MaxArea should be 100 (Square)
	if math.Abs(stats.MaxArea-100.0) > 0.001 {
		t.Errorf("Expected max area 100.0, got %.2f", stats.MaxArea)
	}

	// AvgArea should be total / 3
	expectedAvg := expectedTotal / 3
	if math.Abs(stats.AvgArea-expectedAvg) > 0.001 {
		t.Errorf("Expected avg area %.2f, got %.2f", expectedAvg, stats.AvgArea)
	}
}

// TestCompareShapes verifies shape comparison.
func TestCompareShapes(t *testing.T) {
	small := Circle{Radius: 3} // Area ≈ 28.27
	large := Square{Side: 10}  // Area = 100

	result := CompareShapes(small, large)
	if result != -1 {
		t.Errorf("Expected -1 (small < large), got %d", result)
	}

	result = CompareShapes(large, small)
	if result != 1 {
		t.Errorf("Expected 1 (large > small), got %d", result)
	}

	result = CompareShapes(small, small)
	if result != 0 {
		t.Errorf("Expected 0 (equal), got %d", result)
	}
}
