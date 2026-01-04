package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSum(t *testing.T) {
	assert.Equal(t, 8, CalculateSum(3, 5))
	assert.Equal(t, 0, CalculateSum(-5, 5))
	assert.Equal(t, -10, CalculateSum(-5, -5))
}

func TestSwapValues(t *testing.T) {
	a, b := SwapValues(10, 20)
	assert.Equal(t, 20, a)
	assert.Equal(t, 10, b)

	c, d := SwapValues(100, -50)
	assert.Equal(t, -50, c)
	assert.Equal(t, 100, d)
}

func TestIsEven(t *testing.T) {
	assert.True(t, IsEven(4))
	assert.True(t, IsEven(0))
	assert.True(t, IsEven(-2))
	assert.False(t, IsEven(7))
	assert.False(t, IsEven(-3))
}

func TestGetGrade(t *testing.T) {
	assert.Equal(t, "A", GetGrade(95))
	assert.Equal(t, "A", GetGrade(90))
	assert.Equal(t, "B", GetGrade(85))
	assert.Equal(t, "C", GetGrade(75))
	assert.Equal(t, "D", GetGrade(65))
	assert.Equal(t, "D", GetGrade(60))
	assert.Equal(t, "F", GetGrade(55))
}

func TestFindMax(t *testing.T) {
	assert.Equal(t, 10, FindMax([]int{1, 5, 10, 3}))
	assert.Equal(t, -2, FindMax([]int{-5, -2, -10}))
	assert.Equal(t, 0, FindMax([]int{0, -1, -5}))
	assert.Equal(t, 0, FindMax([]int{}))
}

func TestCountVowels(t *testing.T) {
	assert.Equal(t, 2, CountVowels("hello"))
	assert.Equal(t, 5, CountVowels("beautiful")) // b-e-a-u-t-i-f-u-l has 5 vowels: e, a, u, i, u
	assert.Equal(t, 5, CountVowels("aeiou"))
	assert.Equal(t, 0, CountVowels("xyz"))
	assert.Equal(t, 0, CountVowels(""))
}

func TestReverseSlice(t *testing.T) {
	assert.Equal(t, []int{5, 4, 3, 2, 1}, ReverseSlice([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, []int{3, 2, 1}, ReverseSlice([]int{1, 2, 3}))
	assert.Equal(t, []int{1}, ReverseSlice([]int{1}))
	assert.Equal(t, []int{}, ReverseSlice([]int{}))
}

func TestFilterEvens(t *testing.T) {
	assert.Equal(t, []int{2, 4, 6}, FilterEvens([]int{1, 2, 3, 4, 5, 6}))
	assert.Equal(t, []int{}, FilterEvens([]int{1, 3, 5}))
	assert.Equal(t, []int{0, 2, 4}, FilterEvens([]int{0, 2, 4}))
}

func TestMergeMaps(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	result := MergeMaps(m1, m2)

	assert.Equal(t, 1, result["a"])
	assert.Equal(t, 3, result["b"]) // map2 should override
	assert.Equal(t, 4, result["c"])
	assert.Equal(t, 3, len(result))
}

func TestFibonacci(t *testing.T) {
	assert.Equal(t, 0, Fibonacci(0))
	assert.Equal(t, 1, Fibonacci(1))
	assert.Equal(t, 1, Fibonacci(2))
	assert.Equal(t, 2, Fibonacci(3))
	assert.Equal(t, 3, Fibonacci(4))
	assert.Equal(t, 5, Fibonacci(5))
	assert.Equal(t, 8, Fibonacci(6))
	assert.Equal(t, 21, Fibonacci(8))
}

func TestAlternativeFibonacciRecursive(t *testing.T) {
	assert.Equal(t, 0, AlternativeFibonacciRecursive(0))
	assert.Equal(t, 1, AlternativeFibonacciRecursive(1))
	assert.Equal(t, 8, AlternativeFibonacciRecursive(6))
}
