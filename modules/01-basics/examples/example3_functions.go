package examples

import "fmt"

// FunctionExamples demonstrates Go functions.

// BasicFunction is a simple function with parameters and return value.
func BasicFunction(a, b int) int {
	return a + b
}

// MultipleReturns demonstrates returning multiple values (idiomatic for errors).
func MultipleReturns(a, b int) (int, int) {
	return a + b, a * b
}

// NamedReturns demonstrates named return values.
func NamedReturns(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return // naked return
}

// ErrorHandling demonstrates idiomatic error handling.
func ErrorHandling(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// VariadicFunction accepts variable number of arguments.
func VariadicFunction(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// HigherOrderFunction takes a function as parameter.
func HigherOrderFunction(f func(int) int, value int) int {
	return f(value)
}

// ReturnsFunction returns a function.
func ReturnsFunction(multiplier int) func(int) int {
	return func(x int) int {
		return x * multiplier
	}
}

// Closures demonstrates closures capturing variables.
func Closures() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

// DemonstrateBasicFunction shows basic function usage.
func DemonstrateBasicFunction() {
	result := BasicFunction(3, 5)
	fmt.Printf("3 + 5 = %d\n", result)
}

// DemonstrateMultipleReturns shows multiple return values.
func DemonstrateMultipleReturns() {
	sum, product := MultipleReturns(3, 5)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)

	// Ignore one return value
	s, _ := MultipleReturns(10, 20)
	fmt.Printf("Sum only: %d\n", s)
}

// DemonstrateNamedReturns shows named return values.
func DemonstrateNamedReturns() {
	sum, product := NamedReturns(4, 7)
	fmt.Printf("Named returns - Sum: %d, Product: %d\n", sum, product)
}

// DemonstrateErrorHandling shows error handling pattern.
func DemonstrateErrorHandling() {
	// Successful case
	result, err := ErrorHandling(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("10 / 2 = %f\n", result)

	// Error case
	_, err = ErrorHandling(10, 0)
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}
}

// DemonstrateVariadicFunction shows variadic functions.
func DemonstrateVariadicFunction() {
	// Variable number of arguments
	sum1 := VariadicFunction(1, 2, 3)
	sum2 := VariadicFunction(1, 2, 3, 4, 5)
	sum3 := VariadicFunction()

	fmt.Printf("Sum of 1,2,3: %d\n", sum1)
	fmt.Printf("Sum of 1,2,3,4,5: %d\n", sum2)
	fmt.Printf("Sum of nothing: %d\n", sum3)

	// Spread slice as arguments
	numbers := []int{10, 20, 30}
	sum4 := VariadicFunction(numbers...)
	fmt.Printf("Sum of slice: %d\n", sum4)
}

// DemonstrateHigherOrderFunction shows functions as parameters.
func DemonstrateHigherOrderFunction() {
	double := func(x int) int {
		return x * 2
	}

	square := func(x int) int {
		return x * x
	}

	result1 := HigherOrderFunction(double, 5)
	result2 := HigherOrderFunction(square, 5)

	fmt.Printf("Double 5: %d\n", result1)
	fmt.Printf("Square 5: %d\n", result2)
}

// DemonstrateReturnsFunction shows functions returning functions.
func DemonstrateReturnsFunction() {
	double := ReturnsFunction(2)
	triple := ReturnsFunction(3)

	fmt.Printf("Double 5: %d\n", double(5))
	fmt.Printf("Triple 5: %d\n", triple(5))
}

// DemonstrateClosures shows closures.
func DemonstrateClosures() {
	counter1 := Closures()
	counter2 := Closures()

	fmt.Printf("Counter1: %d\n", counter1()) // 1
	fmt.Printf("Counter1: %d\n", counter1()) // 2
	fmt.Printf("Counter2: %d\n", counter2()) // 1
	fmt.Printf("Counter1: %d\n", counter1()) // 3
}
