package examples

import "fmt"

// ControlFlowExamples demonstrates Go's control flow structures.

// IfStatements demonstrates if statements in Go.
func IfStatements() {
	age := 25

	// Basic if
	if age >= 18 {
		fmt.Println("Adult")
	}

	// If-else
	if age >= 65 {
		fmt.Println("Senior")
	} else if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	// If with initialization (scope is limited to if block)
	if doubled := age * 2; doubled > 40 {
		fmt.Printf("Doubled age %d is over 40\n", doubled)
	}
	// doubled is not available here
}

// ForLoops demonstrates for loop variations in Go.
func ForLoops() {
	// Traditional for loop
	fmt.Println("Count to 5:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// While-style for loop
	fmt.Println("While-style countdown:")
	count := 5
	for count > 0 {
		fmt.Printf("%d ", count)
		count--
	}
	fmt.Println()

	// Infinite loop with break
	fmt.Println("Infinite loop with break:")
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()

	// Continue statement
	fmt.Println("Skip even numbers:")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// RangeLoops demonstrates range loops over different types.
func RangeLoops() {
	// Range over slice
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Range over slice:")
	for index, value := range numbers {
		fmt.Printf("Index %d: Value %d\n", index, value)
	}

	// Range with only index
	fmt.Println("Only indices:")
	for index := range numbers {
		fmt.Printf("%d ", index)
	}
	fmt.Println()

	// Range with only value (use _ for unwanted values)
	fmt.Println("Only values:")
	for _, value := range numbers {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

	// Range over map
	scores := map[string]int{"Alice": 100, "Bob": 85, "Carol": 92}
	fmt.Println("Range over map:")
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}

	// Range over string (iterates over runes, not bytes!)
	text := "Hello, 世界"
	fmt.Println("Range over string:")
	for index, char := range text {
		fmt.Printf("Index %d: %c\n", index, char)
	}
}

// SwitchStatements demonstrates switch statements in Go.
func SwitchStatements() {
	day := "Monday"

	// Basic switch (no fallthrough by default!)
	switch day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Friday":
		fmt.Println("Almost weekend!")
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Midweek")
	}

	// Switch with initialization
	switch hour := 14; {
	case hour < 12:
		fmt.Println("Morning")
	case hour < 17:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Evening")
	}

	// Switch without condition (like if-else chain)
	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature < 15:
		fmt.Println("Cold")
	case temperature < 25:
		fmt.Println("Mild")
	default:
		fmt.Println("Warm")
	}

	// Type switch (we'll cover this more in interfaces module)
	var value interface{} = 42
	switch v := value.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	default:
		fmt.Printf("Unknown type\n")
	}
}

// DeferStatement demonstrates the defer keyword.
func DeferStatement() {
	fmt.Println("Start")

	// Defer executes when function returns
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	defer fmt.Println("Deferred 3")

	fmt.Println("Middle")
	fmt.Println("End")

	// Output order:
	// Start
	// Middle
	// End
	// Deferred 3 (LIFO order!)
	// Deferred 2
	// Deferred 1
}

// DeferWithArguments demonstrates that deferred function arguments are evaluated immediately.
func DeferWithArguments() {
	x := 10

	defer fmt.Println("Deferred x:", x) // x evaluated now (10)

	x = 20
	fmt.Println("Current x:", x) // 20

	// Output:
	// Current x: 20
	// Deferred x: 10
}
