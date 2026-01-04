// Package examples demonstrates basic Go concepts for experienced developers.
//
// This file shows:
// - Package structure and imports
// - Variable declarations and zero values
// - Basic types and type inference
// - Constants
package examples

import "fmt"

// Variables demonstrates different ways to declare variables in Go.
func Variables() {
	// Explicit type declaration
	var name string
	var age int
	var isActive bool

	fmt.Printf("Zero values - name: '%s', age: %d, isActive: %v\n", name, age, isActive)

	// Initialization with values
	var city string = "San Francisco"
	var population int = 800000

	fmt.Printf("Initialized - city: %s, population: %d\n", city, population)

	// Type inference (var without type)
	var country = "USA"
	var zipCode = 94102

	fmt.Printf("Inferred - country: %s, zipCode: %d\n", country, zipCode)

	// Short variable declaration (only in functions)
	language := "Go"
	version := 1.21

	fmt.Printf("Short declaration - language: %s, version: %.2f\n", language, version)

	// Multiple variable declaration
	var (
		firstName = "Alice"
		lastName  = "Smith"
		userAge   = 30
	)

	fmt.Printf("Multiple declaration - %s %s, age %d\n", firstName, lastName, userAge)
}

// Constants demonstrates constant declarations.
func Constants() {
	// Simple constants
	const Pi = 3.14159
	const MaxRetries = 3
	const AppName = "LearningGo"

	fmt.Printf("Constants - Pi: %f, MaxRetries: %d, AppName: %s\n", Pi, MaxRetries, AppName)

	// Constant block
	const (
		StatusOK    = 200
		StatusError = 500
	)

	fmt.Printf("Status codes - OK: %d, Error: %d\n", StatusOK, StatusError)

	// iota - automatic incrementing constant
	const (
		Monday = iota // 0
		Tuesday       // 1
		Wednesday     // 2
		Thursday      // 3
		Friday        // 4
	)

	fmt.Printf("Days - Monday: %d, Friday: %d\n", Monday, Friday)
}

// BasicTypes demonstrates Go's basic types.
func BasicTypes() {
	// Integer types
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var u8 uint8 = 255

	// int is platform-dependent (32 or 64 bit)
	var i int = 42

	fmt.Printf("Integers - int8: %d, int: %d, uint8: %d\n", i8, i, u8)

	// Floating point
	var f32 float32 = 3.14
	var f64 float64 = 3.14159265359

	fmt.Printf("Floats - float32: %f, float64: %f\n", f32, f64)

	// Boolean
	var isTrue bool = true
	var isFalse bool = false

	fmt.Printf("Booleans - true: %v, false: %v\n", isTrue, isFalse)

	// String
	var greeting string = "Hello, Go!"
	var multiline = `This is a
	multi-line string
	using backticks`

	fmt.Printf("Strings - greeting: %s\n", greeting)
	fmt.Printf("Multiline: %s\n", multiline)

	// Rune (alias for int32, represents a Unicode code point)
	var char rune = 'A'
	var emoji rune = '🚀'

	fmt.Printf("Runes - char: %c (%d), emoji: %c (%d)\n", char, char, emoji, emoji)

	// Byte (alias for uint8)
	var b byte = 65 // ASCII 'A'

	fmt.Printf("Byte - b: %c (%d)\n", b, b)

	// Type aliases are useful
	fmt.Printf("i16: %d, i32: %d, i64: %d\n", i16, i32, i64)
}

// Pointers demonstrates pointer usage in Go.
func Pointers() {
	// Basic pointer
	var x int = 42
	var p *int = &x

	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", p)
	fmt.Printf("Value at address (dereferencing): %d\n", *p)

	// Modify through pointer
	*p = 100
	fmt.Printf("After modification - x: %d\n", x)

	// Zero value of pointer is nil
	var nilPointer *int
	fmt.Printf("Nil pointer: %v\n", nilPointer)

	// Creating pointer with new
	ptr := new(int)
	*ptr = 50
	fmt.Printf("Pointer created with new: %d\n", *ptr)
}

// Arrays demonstrates fixed-size arrays in Go.
func Arrays() {
	// Array declaration
	var numbers [5]int
	fmt.Printf("Zero-value array: %v\n", numbers)

	// Array initialization
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Printf("Initialized array: %v\n", primes)

	// Let compiler count
	fibonacci := [...]int{0, 1, 1, 2, 3, 5, 8}
	fmt.Printf("Auto-sized array: %v, length: %d\n", fibonacci, len(fibonacci))

	// Access elements
	fmt.Printf("First prime: %d, Last prime: %d\n", primes[0], primes[4])

	// Arrays are values (copied when assigned)
	copy := primes
	copy[0] = 100
	fmt.Printf("Original: %v, Copy: %v\n", primes, copy)
}

// Slices demonstrates dynamic slices (more common than arrays).
func Slices() {
	// Slice declaration (no size)
	var numbers []int
	fmt.Printf("Nil slice: %v, len: %d, cap: %d\n", numbers, len(numbers), cap(numbers))

	// Create slice with make
	scores := make([]int, 5)       // length 5, capacity 5
	buffer := make([]int, 0, 10)   // length 0, capacity 10
	fmt.Printf("Scores: %v, Buffer: %v\n", scores, buffer)

	// Slice literal
	languages := []string{"Go", "Python", "Rust", "JavaScript"}
	fmt.Printf("Languages: %v\n", languages)

	// Append to slice
	languages = append(languages, "TypeScript")
	fmt.Printf("After append: %v\n", languages)

	// Slicing a slice
	subset := languages[1:3] // From index 1 to 2 (3 is exclusive)
	fmt.Printf("Subset: %v\n", subset)

	// Slices are references
	subset[0] = "Java"
	fmt.Printf("Original after modifying subset: %v\n", languages)
}

// Maps demonstrates Go's hash map type.
func Maps() {
	// Map declaration
	var ages map[string]int
	fmt.Printf("Nil map: %v\n", ages)

	// Create map with make
	ages = make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25
	fmt.Printf("Ages: %v\n", ages)

	// Map literal
	scores := map[string]int{
		"Alice": 100,
		"Bob":   85,
		"Carol": 92,
	}
	fmt.Printf("Scores: %v\n", scores)

	// Access value
	aliceScore := scores["Alice"]
	fmt.Printf("Alice's score: %d\n", aliceScore)

	// Check if key exists
	bobScore, exists := scores["Bob"]
	fmt.Printf("Bob's score: %d, exists: %v\n", bobScore, exists)

	unknownScore, exists := scores["Unknown"]
	fmt.Printf("Unknown score: %d, exists: %v\n", unknownScore, exists)

	// Delete key
	delete(scores, "Bob")
	fmt.Printf("After deleting Bob: %v\n", scores)
}
