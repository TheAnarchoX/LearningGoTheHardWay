package solutions

// SOLUTION: This file contains the corrected versions of the exercises.

import (
	"fmt"
	"strings"
)

// CalculateSum returns the sum of two integers.
func CalculateSum(a, b int) int {
	return a + b // Fixed: Addition instead of subtraction
}

// SwapValues swaps two integer values and returns them.
func SwapValues(a, b int) (int, int) {
	return b, a // Fixed: Swapped return values
}

// IsEven returns true if the number is even, false otherwise.
func IsEven(n int) bool {
	return n%2 == 0 // Fixed: Check for even (remainder 0)
}

// GetGrade returns a letter grade based on a numeric score.
func GetGrade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 { // Fixed: Correct threshold for D
		return "D"
	}
	return "F"
}

// FindMax returns the maximum value in a slice.
func FindMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0] // Fixed: Initialize with first element to handle negatives
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}

// CountVowels counts the number of vowels (a, e, i, o, u) in a string.
func CountVowels(s string) int {
	count := 0
	// Convert to lowercase for case-insensitive comparison
	lower := strings.ToLower(s)

	for _, char := range lower {
		// Fixed: Check all vowels
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			count++
		}
	}
	return count
}

// ReverseSlice reverses a slice in place and returns it.
func ReverseSlice(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - 1 - i
		// Fixed: Proper swap using multiple assignment
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

// FilterEvens returns a new slice containing only even numbers.
func FilterEvens(numbers []int) []int {
	result := []int{}
	for _, n := range numbers {
		if n%2 == 0 { // Fixed: Check for even numbers
			result = append(result, n)
		}
	}
	return result
}

// MergeMaps merges two maps. If a key exists in both, use the value from map2.
func MergeMaps(map1, map2 map[string]int) map[string]int {
	result := make(map[string]int)

	// Copy map1
	for k, v := range map1 {
		result[k] = v
	}

	// Fixed: Copy map2, overriding any duplicate keys
	for k, v := range map2 {
		result[k] = v
	}

	return result
}

// Fibonacci returns the nth Fibonacci number.
// The sequence is: 0, 1, 1, 2, 3, 5, 8, 13, 21, ...
func Fibonacci(n int) int {
	// Base cases
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// Iterative approach (more efficient than recursive)
	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

// AlternativeFibonacciRecursive is a recursive implementation (simpler but less efficient).
func AlternativeFibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return AlternativeFibonacciRecursive(n-1) + AlternativeFibonacciRecursive(n-2)
}

// DemonstrateSolutions runs all fixed functions for demonstration.
func DemonstrateSolutions() {
	fmt.Println("Running fixed functions:")

	fmt.Printf("Sum of 3 and 5: %d\n", CalculateSum(3, 5))

	a, b := SwapValues(10, 20)
	fmt.Printf("Swapped 10 and 20: %d, %d\n", a, b)

	fmt.Printf("Is 4 even? %v\n", IsEven(4))
	fmt.Printf("Is 7 even? %v\n", IsEven(7))

	fmt.Printf("Grade for 85: %s\n", GetGrade(85))
	fmt.Printf("Grade for 65: %s\n", GetGrade(65))

	fmt.Printf("Max of [-5, -2, -10]: %d\n", FindMax([]int{-5, -2, -10}))

	fmt.Printf("Vowels in 'hello': %d\n", CountVowels("hello"))

	reversed := ReverseSlice([]int{1, 2, 3, 4, 5})
	fmt.Printf("Reversed [1,2,3,4,5]: %v\n", reversed)

	evens := FilterEvens([]int{1, 2, 3, 4, 5, 6})
	fmt.Printf("Evens from [1,2,3,4,5,6]: %v\n", evens)

	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	merged := MergeMaps(m1, m2)
	fmt.Printf("Merged maps: %v\n", merged)

	fmt.Printf("Fibonacci(6): %d\n", Fibonacci(6))
	fmt.Printf("Fibonacci(8): %d\n", Fibonacci(8))
}
