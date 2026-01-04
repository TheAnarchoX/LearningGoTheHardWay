package examples

import (
	"math"
	"testing"
)

// TestInterfaceBasics verifies interface usage examples.
func TestInterfaceBasics(t *testing.T) {
	InterfaceBasics()
}

// TestCommonInterfaces verifies standard library interface examples.
func TestCommonInterfaces(t *testing.T) {
	CommonInterfaces()
}

// TestInterfaceComposition verifies interface embedding examples.
func TestInterfaceComposition(t *testing.T) {
	InterfaceComposition()
}

// TestTypeAssertion verifies type assertion examples.
func TestTypeAssertion(t *testing.T) {
	TypeAssertion()
}

// TestTypeSwitch verifies type switch examples.
func TestTypeSwitch(t *testing.T) {
	TypeSwitch()
}

// TestEmptyInterface verifies empty interface examples.
func TestEmptyInterface(t *testing.T) {
	EmptyInterface()
}

// TestEmptyInterfaceWithTypeSwitch verifies type switch with empty interface.
func TestEmptyInterfaceWithTypeSwitch(t *testing.T) {
	EmptyInterfaceWithTypeSwitch()
}

// TestNilInterface verifies nil interface behavior.
func TestNilInterface(t *testing.T) {
	NilInterface()
}

// TestPolymorphism verifies polymorphic behavior.
func TestPolymorphism(t *testing.T) {
	Polymorphism()
}

// TestCircleImplementsShape verifies Circle implements Shape.
func TestCircleImplementsShape(t *testing.T) {
	circle := Circle{Radius: 5}

	area := circle.Area()
	expectedArea := math.Pi * 25
	if math.Abs(area-expectedArea) > 0.001 {
		t.Errorf("Expected area %.2f, got %.2f", expectedArea, area)
	}

	perimeter := circle.Perimeter()
	expectedPerimeter := 2 * math.Pi * 5
	if math.Abs(perimeter-expectedPerimeter) > 0.001 {
		t.Errorf("Expected perimeter %.2f, got %.2f", expectedPerimeter, perimeter)
	}

	// Verify Circle can be assigned to Shape interface
	var s Shape = circle
	if s == nil {
		t.Error("Circle should implement Shape interface")
	}
}

// TestRectangleImplementsShape verifies Rectangle implements Shape.
func TestRectangleImplementsShape(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 5}

	area := rect.Area()
	expectedArea := 50.0
	if area != expectedArea {
		t.Errorf("Expected area %.2f, got %.2f", expectedArea, area)
	}

	perimeter := rect.Perimeter()
	expectedPerimeter := 30.0
	if perimeter != expectedPerimeter {
		t.Errorf("Expected perimeter %.2f, got %.2f", expectedPerimeter, perimeter)
	}

	// Verify Rectangle can be assigned to Shape interface
	var s Shape = rect
	if s == nil {
		t.Error("Rectangle should implement Shape interface")
	}
}

// TestTriangleImplementsShape verifies Triangle implements Shape.
func TestTriangleImplementsShape(t *testing.T) {
	triangle := Triangle{Base: 10, Height: 8, SideA: 6, SideB: 8, SideC: 10}

	area := triangle.Area()
	expectedArea := 40.0
	if area != expectedArea {
		t.Errorf("Expected area %.2f, got %.2f", expectedArea, area)
	}

	perimeter := triangle.Perimeter()
	expectedPerimeter := 24.0
	if perimeter != expectedPerimeter {
		t.Errorf("Expected perimeter %.2f, got %.2f", expectedPerimeter, perimeter)
	}

	// Verify Triangle can be assigned to Shape interface
	var s Shape = triangle
	if s == nil {
		t.Error("Triangle should implement Shape interface")
	}
}

// TestCircleString verifies Circle implements Stringer.
func TestCircleString(t *testing.T) {
	circle := Circle{Radius: 5}
	str := circle.String()
	expected := "Circle(radius=5.00)"
	if str != expected {
		t.Errorf("Expected %q, got %q", expected, str)
	}
}

// TestRectangleString verifies Rectangle implements Stringer.
func TestRectangleString(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 5}
	str := rect.String()
	expected := "Rectangle(width=10.00, height=5.00)"
	if str != expected {
		t.Errorf("Expected %q, got %q", expected, str)
	}
}

// TestTypeAssertionSuccess verifies successful type assertion.
func TestTypeAssertionSuccess(t *testing.T) {
	var s Shape = Circle{Radius: 5}

	circle, ok := s.(Circle)
	if !ok {
		t.Error("Type assertion should succeed for Circle")
	}

	if circle.Radius != 5 {
		t.Errorf("Expected radius 5, got %.2f", circle.Radius)
	}
}

// TestTypeAssertionFailure verifies failed type assertion.
func TestTypeAssertionFailure(t *testing.T) {
	var s Shape = Circle{Radius: 5}

	_, ok := s.(Rectangle)
	if ok {
		t.Error("Type assertion should fail for Rectangle when Shape is Circle")
	}
}

// TestProcessValue verifies ProcessValue with different types.
func TestProcessValue(t *testing.T) {
	// Just ensure it doesn't panic
	ProcessValue(42)
	ProcessValue("test")
	ProcessValue(Circle{Radius: 5})
	ProcessValue([]int{1, 2, 3})
	ProcessValue(true)
}

// TestShapeSlice verifies slice of Shape interface.
func TestShapeSlice(t *testing.T) {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 10, Height: 5},
		Triangle{Base: 10, Height: 8, SideA: 6, SideB: 8, SideC: 10},
	}

	if len(shapes) != 3 {
		t.Errorf("Expected 3 shapes, got %d", len(shapes))
	}

	// Verify all shapes can calculate area
	for i, shape := range shapes {
		area := shape.Area()
		if area <= 0 {
			t.Errorf("Shape %d has invalid area: %.2f", i, area)
		}
	}
}
