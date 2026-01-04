package exercises

// EXERCISE: Fix the bugs in this file to make the tests pass.
// Each function has intentional bugs marked with // BUG: comments.

import "fmt"

// CalculateSum should return the sum of two integers.
// BUG: This function has a logic error.
func CalculateSum(a, b int) int {
	return a - b // BUG: Should be addition, not subtraction
}

// SwapValues should swap two integer values and return them.
// BUG: The values are not actually being swapped.
func SwapValues(a, b int) (int, int) {
	return a, b // BUG: Should return b, a
}

// IsEven should return true if the number is even, false otherwise.
// BUG: The logic is inverted.
func IsEven(n int) bool {
	return n%2 != 0 // BUG: Should be == 0
}

// GetGrade should return a letter grade based on a numeric score.
// Score >= 90: "A"
// Score >= 80: "B"
// Score >= 70: "C"
// Score >= 60: "D"
// Score < 60: "F"
// BUG: The conditions are wrong.
func GetGrade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 50 { // BUG: Should be >= 60
		return "D"
	}
	return "F"
}

// FindMax should return the maximum value in a slice.
// BUG: The initial max value is wrong.
func FindMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := 0 // BUG: Should be numbers[0] to handle negative numbers
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}

// CountVowels should count the number of vowels (a, e, i, o, u) in a string.
// BUG: Missing vowels in the check.
func CountVowels(s string) int {
	count := 0
	for _, char := range s {
		lower := string(char)
		// BUG: Missing 'i', 'o', 'u'
		if lower == "a" || lower == "e" {
			count++
		}
	}
	return count
}

// ReverseSlice should reverse a slice in place and return it.
// BUG: The swap logic is incomplete.
func ReverseSlice(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - 1 - i
		// BUG: Missing actual swap
		numbers[i] = numbers[j]
		// Should be:
		// numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

// FilterEvens should return a new slice containing only even numbers.
// BUG: Filtering odds instead of evens.
func FilterEvens(numbers []int) []int {
	result := []int{}
	for _, n := range numbers {
		if n%2 != 0 { // BUG: Should be == 0
			result = append(result, n)
		}
	}
	return result
}

// MergeMaps should merge two maps. If a key exists in both, use the value from map2.
// BUG: Not handling the merge correctly.
func MergeMaps(map1, map2 map[string]int) map[string]int {
	result := make(map[string]int)

	// Copy map1
	for k, v := range map1 {
		result[k] = v
	}

	// BUG: map2 values should override map1 values, but we're not copying map2
	// Missing: copy map2 to result

	return result
}

// TODO: Implement this function
// Fibonacci should return the nth Fibonacci number.
// The sequence is: 0, 1, 1, 2, 3, 5, 8, 13, 21, ...
// F(0) = 0, F(1) = 1, F(n) = F(n-1) + F(n-2)
func Fibonacci(n int) int {
	// TODO: Implement the Fibonacci function
	return 0 // Placeholder
}

// DemonstrateBugs runs all buggy functions for demonstration.
func DemonstrateBugs() {
	fmt.Println("Running buggy functions (fix these!):")

	fmt.Printf("Sum of 3 and 5: %d (should be 8)\n", CalculateSum(3, 5))

	a, b := SwapValues(10, 20)
	fmt.Printf("Swapped 10 and 20: %d, %d (should be 20, 10)\n", a, b)

	fmt.Printf("Is 4 even? %v (should be true)\n", IsEven(4))
	fmt.Printf("Is 7 even? %v (should be false)\n", IsEven(7))

	fmt.Printf("Grade for 85: %s (should be B)\n", GetGrade(85))
	fmt.Printf("Grade for 65: %s (should be D)\n", GetGrade(65))

	fmt.Printf("Max of [-5, -2, -10]: %d (should be -2)\n", FindMax([]int{-5, -2, -10}))

	fmt.Printf("Vowels in 'hello': %d (should be 2)\n", CountVowels("hello"))

	reversed := ReverseSlice([]int{1, 2, 3, 4, 5})
	fmt.Printf("Reversed [1,2,3,4,5]: %v (should be [5,4,3,2,1])\n", reversed)

	evens := FilterEvens([]int{1, 2, 3, 4, 5, 6})
	fmt.Printf("Evens from [1,2,3,4,5,6]: %v (should be [2,4,6])\n", evens)

	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	merged := MergeMaps(m1, m2)
	fmt.Printf("Merged maps: %v (should be map[a:1 b:3 c:4])\n", merged)

	fmt.Printf("Fibonacci(6): %d (should be 8)\n", Fibonacci(6))
}
