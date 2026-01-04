// Package examples demonstrates Go's type system for experienced developers.
//
// This file shows:
// - Struct definition and initialization
// - Methods with value and pointer receivers
// - Constructor patterns
// - Struct tags for JSON, XML, etc.
package examples

import (
	"encoding/json"
	"fmt"
)

// Person represents a basic struct with fields.
//
// Coming from Java/C++: No classes, just structs with methods defined separately.
// Coming from Python: Similar to dataclasses but with explicit types.
// Coming from JavaScript: Like objects but with fixed structure and types.
type Person struct {
	Name string
	Age  int
	City string
}

// Greet demonstrates a method with a value receiver.
// Value receivers operate on a copy of the struct.
//
// Coming from OOP: Methods are defined outside the struct, not inside.
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s from %s", p.Name, p.City)
}

// HaveBirthday demonstrates a method with a pointer receiver.
// Pointer receivers can modify the original struct.
//
// Rule of thumb: Use pointer receivers when you need to modify the struct
// or when the struct is large (to avoid copying).
func (p *Person) HaveBirthday() {
	p.Age++
}

// IsAdult demonstrates a value receiver for read-only operations.
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// StructBasics demonstrates struct creation and initialization.
func StructBasics() {
	fmt.Println("=== Struct Basics ===")

	// Zero value initialization
	var p1 Person
	fmt.Printf("Zero value: %+v\n", p1) // %+v shows field names

	// Struct literal with field names (recommended)
	p2 := Person{
		Name: "Alice",
		Age:  30,
		City: "San Francisco",
	}
	fmt.Printf("Named fields: %+v\n", p2)

	// Struct literal without field names (not recommended - fragile)
	p3 := Person{"Bob", 25, "New York"}
	fmt.Printf("Positional: %+v\n", p3)

	// Partial initialization (unspecified fields get zero values)
	p4 := Person{
		Name: "Charlie",
		Age:  35,
	}
	fmt.Printf("Partial: %+v\n", p4)

	// Using new - returns a pointer to zero-valued struct
	p5 := new(Person)
	p5.Name = "Diana"
	p5.Age = 28
	fmt.Printf("With new: %+v\n", *p5)
}

// MethodReceivers demonstrates value vs pointer receivers.
func MethodReceivers() {
	fmt.Println("\n=== Method Receivers ===")

	alice := Person{Name: "Alice", Age: 30, City: "NYC"}

	// Value receiver - operates on a copy
	greeting := alice.Greet()
	fmt.Println(greeting)

	// Pointer receiver - modifies original
	fmt.Printf("Before birthday: Age = %d\n", alice.Age)
	alice.HaveBirthday()
	fmt.Printf("After birthday: Age = %d\n", alice.Age)

	// Go automatically handles pointer dereferencing
	ptrAlice := &Person{Name: "Bob", Age: 20, City: "SF"}
	ptrAlice.HaveBirthday() // Works on pointer
	fmt.Printf("Pointer birthday: Age = %d\n", ptrAlice.Age)

	// Value receiver also works with pointer (Go dereferences automatically)
	fmt.Println(ptrAlice.Greet())
}

// Employee demonstrates struct embedding (composition).
//
// Coming from Java/C++: This is composition, not inheritance.
// Coming from Python: Similar to composition but with field promotion.
type Employee struct {
	Person     // Embedded struct - fields are "promoted"
	EmployeeID int
	Department string
}

// NewEmployee is a constructor function pattern.
// Go doesn't have constructors, but this pattern is idiomatic.
func NewEmployee(name, city, department string, age, employeeID int) *Employee {
	return &Employee{
		Person: Person{
			Name: name,
			Age:  age,
			City: city,
		},
		EmployeeID: employeeID,
		Department: department,
	}
}

// Work demonstrates a method specific to Employee.
func (e Employee) Work() string {
	return fmt.Sprintf("%s is working in %s department", e.Name, e.Department)
}

// Embedding demonstrates struct composition.
func Embedding() {
	fmt.Println("\n=== Struct Embedding ===")

	emp := NewEmployee("Alice", "Boston", "Engineering", 30, 1001)

	// Access promoted fields directly
	fmt.Printf("Name: %s, Age: %d\n", emp.Name, emp.Age)

	// Call promoted methods
	fmt.Println(emp.Greet())

	// Call employee-specific methods
	fmt.Println(emp.Work())

	// Can still access embedded struct explicitly
	fmt.Printf("Person: %+v\n", emp.Person)

	// Modify using promoted method
	emp.HaveBirthday()
	fmt.Printf("After birthday: Age = %d\n", emp.Age)
}

// UserAccount demonstrates struct tags for JSON marshaling.
//
// Coming from Java: Like Jackson annotations.
// Coming from Python: Like dataclass field metadata.
// Coming from JavaScript: Enables JSON serialization control.
type UserAccount struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"` // "-" means omit from JSON
	IsActive bool   `json:"is_active,omitempty"`
}

// StructTags demonstrates struct tags for encoding.
func StructTags() {
	fmt.Println("\n=== Struct Tags ===")

	user := UserAccount{
		ID:       1,
		Username: "alice",
		Email:    "alice@example.com",
		Password: "secret123",
		IsActive: true,
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}
	fmt.Printf("JSON: %s\n", jsonData)

	// Password is omitted due to "-" tag
	// Notice the field names use json tags

	// Unmarshal from JSON
	jsonStr := `{"id":2,"username":"bob","email":"bob@example.com"}`
	var user2 UserAccount
	if err := json.Unmarshal([]byte(jsonStr), &user2); err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
		return
	}
	fmt.Printf("Unmarshaled: %+v\n", user2)
}

// Point demonstrates comparable structs.
type Point struct {
	X, Y int
}

// Rect demonstrates struct with struct fields.
type Rect struct {
	TopLeft     Point
	BottomRight Point
}

// Area calculates the area of the rectangle.
func (r Rect) Area() int {
	width := r.BottomRight.X - r.TopLeft.X
	height := r.TopLeft.Y - r.BottomRight.Y
	return width * height
}

// NestedStructs demonstrates structs containing other structs.
func NestedStructs() {
	fmt.Println("\n=== Nested Structs ===")

	rect := Rect{
		TopLeft:     Point{X: 0, Y: 10},
		BottomRight: Point{X: 10, Y: 0},
	}

	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %d\n", rect.Area())

	// Structs are comparable if all fields are comparable
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 1, Y: 2}
	p3 := Point{X: 3, Y: 4}

	fmt.Printf("p1 == p2: %v\n", p1 == p2) // true
	fmt.Printf("p1 == p3: %v\n", p1 == p3) // false
}

// Config demonstrates anonymous structs for one-off data structures.
//
// Coming from JavaScript: Like object literals.
// Coming from Python: Like creating a simple class on the fly.
func AnonymousStructs() {
	fmt.Println("\n=== Anonymous Structs ===")

	// Anonymous struct literal
	config := struct {
		Host string
		Port int
		TLS  bool
	}{
		Host: "localhost",
		Port: 8080,
		TLS:  false,
	}

	fmt.Printf("Config: %+v\n", config)

	// Useful for test data or temporary data structures
	testCases := []struct {
		input    int
		expected int
	}{
		{input: 1, expected: 2},
		{input: 2, expected: 4},
		{input: 3, expected: 6},
	}

	for _, tc := range testCases {
		result := tc.input * 2
		fmt.Printf("Input: %d, Expected: %d, Got: %d\n", tc.input, tc.expected, result)
	}
}
