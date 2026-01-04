// Package solutions contains correct implementations of the exercises.
//
// SOLUTION 2: Type Assertions and Type Switches
//
// This file shows the correct implementation with all bugs fixed.
//
// Key fixes:
// - Stringify() handles all types including nil
// - GetIntValue() uses safe type assertion
// - IsNumeric() has correct logic
// - Sum() handles []float64 case
// - ProcessValue() uses safe type assertions
// - GetTypeName() handles all common types
// - ConvertToInt() returns proper errors
// - Filter() uses safe type assertions
// - Max() handles type checking correctly
// - TypeSafeMap methods use safe type assertions
package solutions

import (
	"fmt"
	"strconv"
)

// Stringify converts any value to a string.
// FIXED: Handles all types including nil
func Stringify(v interface{}) string {
	// FIXED: Handle nil case
	if v == nil {
		return "<nil>"
	}

	switch val := v.(type) {
	case string:
		return val
	case int:
		// FIXED: Use strconv.Itoa for proper conversion
		return strconv.Itoa(val)
	case float64:
		return fmt.Sprintf("%.2f", val)
	case bool:
		if val {
			return "true"
		}
		return "false"
	default:
		return "unknown"
	}
}

// GetIntValue extracts an int from an interface{}.
// FIXED: Uses safe type assertion
func GetIntValue(v interface{}) int {
	// FIXED: Safe type assertion with ok check
	if value, ok := v.(int); ok {
		return value
	}
	return 0
}

// IsNumeric checks if a value is numeric (int or float64).
// FIXED: Correct logic
func IsNumeric(v interface{}) bool {
	switch v.(type) {
	case int:
		return true // FIXED: Return true for int
	case float64:
		return true // FIXED: Return true for float64
	default:
		return false // FIXED: Return false for non-numeric
	}
}

// Sum calculates the sum of values in a slice.
// Works with []int, []float64, or []interface{} containing numbers.
// FIXED: Handles all numeric slice types
func Sum(values interface{}) float64 {
	switch v := values.(type) {
	case []int:
		total := 0.0
		for _, num := range v {
			total += float64(num)
		}
		return total
	case []float64:
		// FIXED: Added case for []float64
		total := 0.0
		for _, num := range v {
			total += num
		}
		return total
	case []interface{}:
		total := 0.0
		for _, val := range v {
			// FIXED: Handle both int and float64 in interface slice
			switch num := val.(type) {
			case int:
				total += float64(num)
			case float64:
				total += num
			}
		}
		return total
	default:
		return 0.0
	}
}

// ProcessValue processes a value based on its type.
// FIXED: Uses safe type assertions
func ProcessValue(v interface{}) string {
	// FIXED: Safe type assertion for string
	if str, ok := v.(string); ok {
		return "String: " + str
	}

	// FIXED: Safe type assertion for int
	if num, ok := v.(int); ok {
		return fmt.Sprintf("Integer: %d", num)
	}

	return "Unknown"
}

// GetTypeName returns the type name of a value.
// FIXED: Handles all common types
func GetTypeName(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float64:
		// FIXED: Added case for float64
		return "float64"
	case bool:
		// FIXED: Added case for bool
		return "bool"
	case []int:
		// FIXED: Added case for []int
		return "[]int"
	default:
		return "unknown"
	}
}

// ConvertToInt converts various types to int.
// FIXED: Handles errors properly
func ConvertToInt(v interface{}) (int, error) {
	switch val := v.(type) {
	case int:
		return val, nil
	case float64:
		return int(val), nil
	case string:
		// FIXED: Return error from strconv.Atoi
		result, err := strconv.Atoi(val)
		if err != nil {
			return 0, fmt.Errorf("cannot convert string %q to int: %w", val, err)
		}
		return result, nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	default:
		// FIXED: Return proper error
		return 0, fmt.Errorf("cannot convert type %T to int", v)
	}
}

// Filter filters a slice based on a type predicate.
// Returns only values of the specified type.
// FIXED: Uses safe type assertions
func Filter(values []interface{}, typeName string) []interface{} {
	result := []interface{}{}

	for _, v := range values {
		switch typeName {
		case "int":
			// FIXED: Safe type assertion
			if val, ok := v.(int); ok {
				result = append(result, val)
			}
		case "string":
			// FIXED: Safe type assertion
			if val, ok := v.(string); ok {
				result = append(result, val)
			}
		case "float64":
			// FIXED: Added case for float64
			if val, ok := v.(float64); ok {
				result = append(result, val)
			}
		case "bool":
			// FIXED: Added case for bool
			if val, ok := v.(bool); ok {
				result = append(result, val)
			}
		}
	}

	return result
}

// Max returns the maximum value from a slice of comparable values.
// FIXED: Handles type checking correctly
func Max(values []interface{}) interface{} {
	if len(values) == 0 {
		return nil
	}

	// FIXED: Check first value type safely
	max, ok := values[0].(int)
	if !ok {
		return nil
	}

	for i := 1; i < len(values); i++ {
		// FIXED: Safe type assertion
		if val, ok := values[i].(int); ok {
			if val > max {
				max = val
			}
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
// FIXED: Checks if key exists and uses safe type assertion
func (m *TypeSafeMap) GetString(key string) string {
	// FIXED: Check if key exists
	value, exists := m.data[key]
	if !exists {
		return ""
	}

	// FIXED: Safe type assertion
	if str, ok := value.(string); ok {
		return str
	}
	return ""
}

// GetInt retrieves an int value.
// FIXED: Checks if key exists and uses safe type assertion
func (m *TypeSafeMap) GetInt(key string) int {
	// FIXED: Check if key exists
	value, exists := m.data[key]
	if !exists {
		return 0
	}

	// FIXED: Safe type assertion
	if num, ok := value.(int); ok {
		return num
	}
	return 0
}

// Contains checks if a key exists.
// FIXED: Properly checks key existence
func (m *TypeSafeMap) Contains(key string) bool {
	// FIXED: Use two-value map access to check key existence
	_, exists := m.data[key]
	return exists
}

// Count returns the number of values of a given type.
// FIXED: Uses safe type assertions
func (m *TypeSafeMap) Count(typeName string) int {
	count := 0
	for _, value := range m.data {
		switch typeName {
		case "string":
			// FIXED: Safe type assertion
			if _, ok := value.(string); ok {
				count++
			}
		case "int":
			// FIXED: Safe type assertion
			if _, ok := value.(int); ok {
				count++
			}
		case "float64":
			if _, ok := value.(float64); ok {
				count++
			}
		case "bool":
			if _, ok := value.(bool); ok {
				count++
			}
		}
	}
	return count
}
