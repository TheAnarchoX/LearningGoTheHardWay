// Package exercises contains practice problems with intentional bugs.
//
// EXERCISE 2: Type Assertions and Type Switches
//
// Fix the bugs in this file to make all tests pass.
// The bugs are marked with // BUG: comments.
//
// Learning objectives:
// - Use type assertions safely
// - Implement type switches correctly
// - Work with empty interface (interface{} / any)
// - Handle type assertion failures
package exercises

import (
	"fmt"
	"strconv"
)

// Stringify converts any value to a string.
// BUG: Doesn't handle all types correctly
func Stringify(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int:
		// BUG: Wrong conversion - should use strconv.Itoa
		return string(val)
	case float64:
		return fmt.Sprintf("%.2f", val)
	case bool:
		if val {
			return "true"
		}
		return "false"
	// BUG: Missing case for nil
	default:
		return "unknown"
	}
}

// GetIntValue extracts an int from an interface{}.
// BUG: Doesn't check if type assertion succeeded
func GetIntValue(v interface{}) int {
	// BUG: Unsafe type assertion - will panic if v is not an int
	value := v.(int)
	return value
}

// IsNumeric checks if a value is numeric (int or float64).
// BUG: Logic is incorrect
func IsNumeric(v interface{}) bool {
	switch v.(type) {
	case int:
		return false // BUG: Should return true
	case float64:
		return false // BUG: Should return true
	default:
		return true // BUG: Should return false
	}
}

// Sum calculates the sum of values in a slice.
// Works with []int, []float64, or []interface{} containing numbers.
// BUG: Type switch doesn't handle all cases
func Sum(values interface{}) float64 {
	switch v := values.(type) {
	case []int:
		total := 0.0
		for _, num := range v {
			total += float64(num)
		}
		return total
	// BUG: Missing case for []float64
	case []interface{}:
		total := 0.0
		for _, val := range v {
			// BUG: Doesn't handle non-numeric values
			num := val.(float64)
			total += num
		}
		return total
	default:
		return 0.0
	}
}

// ProcessValue processes a value based on its type.
// BUG: Type assertions are not checked for success
func ProcessValue(v interface{}) string {
	// Try to treat as string
	str := v.(string) // BUG: Unsafe type assertion
	if str != "" {
		return "String: " + str
	}

	// Try to treat as int
	num := v.(int) // BUG: Unsafe type assertion
	if num != 0 {
		return fmt.Sprintf("Integer: %d", num)
	}

	return "Unknown"
}

// GetTypeName returns the type name of a value.
// BUG: Doesn't handle all common types
func GetTypeName(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	// BUG: Missing cases for other types (float64, bool, etc.)
	default:
		return "unknown"
	}
}

// ConvertToInt converts various types to int.
// BUG: Doesn't handle all cases correctly
func ConvertToInt(v interface{}) (int, error) {
	switch val := v.(type) {
	case int:
		return val, nil
	case float64:
		// BUG: Should round or truncate, not just cast
		return int(val), nil
	case string:
		// BUG: Doesn't handle conversion errors
		result, _ := strconv.Atoi(val)
		return result, nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	default:
		// BUG: Should return an error, not 0
		return 0, nil
	}
}

// Filter filters a slice based on a type predicate.
// Returns only values of the specified type.
// BUG: Implementation is incomplete
func Filter(values []interface{}, typeName string) []interface{} {
	result := []interface{}{}

	// BUG: Type checking logic is wrong
	for _, v := range values {
		switch typeName {
		case "int":
			// BUG: This will panic if v is not an int
			result = append(result, v.(int))
		case "string":
			result = append(result, v.(string))
			// BUG: Missing cases for other types
		}
	}

	return result
}

// Max returns the maximum value from a slice of comparable values.
// BUG: Doesn't handle type checking correctly
func Max(values []interface{}) interface{} {
	if len(values) == 0 {
		return nil
	}

	// BUG: Assumes all values are ints without checking
	max := values[0].(int)

	for i := 1; i < len(values); i++ {
		// BUG: Unsafe type assertion
		val := values[i].(int)
		if val > max {
			max = val
		}
	}

	return max
}

// TypeSafeMap is a map that stores values of any type.
type TypeSafeMap struct {
	data map[string]interface{}
}

// NewTypeSafeMap creates a new TypeSafeMap.
func NewTypeSafeMap() *TypeSafeMap {
	return &TypeSafeMap{
		data: make(map[string]interface{}),
	}
}

// Set stores a value.
func (m *TypeSafeMap) Set(key string, value interface{}) {
	m.data[key] = value
}

// GetString retrieves a string value.
// BUG: Doesn't check if key exists or if value is a string
func (m *TypeSafeMap) GetString(key string) string {
	// BUG: Doesn't check if value exists
	value := m.data[key]
	// BUG: Unsafe type assertion
	return value.(string)
}

// GetInt retrieves an int value.
// BUG: Doesn't check if key exists or if value is an int
func (m *TypeSafeMap) GetInt(key string) int {
	return m.data[key].(int)
}

// Contains checks if a key exists.
// BUG: Implementation is wrong
func (m *TypeSafeMap) Contains(key string) bool {
	// BUG: Should check if key exists, not if value is non-nil
	return m.data[key] != nil
}

// Count returns the number of values of a given type.
// BUG: Type switch doesn't work correctly
func (m *TypeSafeMap) Count(typeName string) int {
	count := 0
	for _, value := range m.data {
		switch typeName {
		case "string":
			// BUG: This will panic if value is not a string
			_ = value.(string)
			count++
		case "int":
			_ = value.(int)
			count++
		}
	}
	return count
}
