package examples

import (
	"testing"
)

// TestStructBasics verifies struct initialization examples.
func TestStructBasics(t *testing.T) {
	StructBasics()
}

// TestMethodReceivers verifies value and pointer receiver examples.
func TestMethodReceivers(t *testing.T) {
	MethodReceivers()
}

// TestEmbedding verifies struct composition examples.
func TestEmbedding(t *testing.T) {
	Embedding()
}

// TestStructTags verifies struct tag examples.
func TestStructTags(t *testing.T) {
	StructTags()
}

// TestNestedStructs verifies nested struct examples.
func TestNestedStructs(t *testing.T) {
	NestedStructs()
}

// TestAnonymousStructs verifies anonymous struct examples.
func TestAnonymousStructs(t *testing.T) {
	AnonymousStructs()
}

// TestPerson verifies Person struct methods.
func TestPerson(t *testing.T) {
	p := Person{Name: "Alice", Age: 30, City: "NYC"}

	// Test Greet
	greeting := p.Greet()
	expected := "Hello, I'm Alice from NYC"
	if greeting != expected {
		t.Errorf("Expected %q, got %q", expected, greeting)
	}

	// Test IsAdult
	if !p.IsAdult() {
		t.Error("Expected Alice to be an adult")
	}

	// Test HaveBirthday
	originalAge := p.Age
	p.HaveBirthday()
	if p.Age != originalAge+1 {
		t.Errorf("Expected age %d, got %d", originalAge+1, p.Age)
	}
}

// TestEmployee verifies Employee struct and constructor.
func TestEmployee(t *testing.T) {
	emp := NewEmployee("Bob", "SF", "Engineering", 25, 1001)

	// Verify fields
	if emp.Name != "Bob" {
		t.Errorf("Expected name Bob, got %s", emp.Name)
	}
	if emp.EmployeeID != 1001 {
		t.Errorf("Expected ID 1001, got %d", emp.EmployeeID)
	}

	// Test promoted method
	greeting := emp.Greet()
	if greeting != "Hello, I'm Bob from SF" {
		t.Errorf("Unexpected greeting: %s", greeting)
	}

	// Test employee-specific method
	work := emp.Work()
	if work != "Bob is working in Engineering department" {
		t.Errorf("Unexpected work message: %s", work)
	}
}

// TestRectArea verifies Rect area calculation.
func TestRectArea(t *testing.T) {
	rect := Rect{
		TopLeft:     Point{X: 0, Y: 10},
		BottomRight: Point{X: 10, Y: 0},
	}

	area := rect.Area()
	expected := 100
	if area != expected {
		t.Errorf("Expected area %d, got %d", expected, area)
	}
}

// TestPointComparison verifies struct comparison.
func TestPointComparison(t *testing.T) {
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 1, Y: 2}
	p3 := Point{X: 3, Y: 4}

	if p1 != p2 {
		t.Error("Expected p1 and p2 to be equal")
	}

	if p1 == p3 {
		t.Error("Expected p1 and p3 to be different")
	}
}
