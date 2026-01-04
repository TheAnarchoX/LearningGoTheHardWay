package exercises

import (
	"testing"
)

// TestStringify verifies Stringify function.
func TestStringify(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{"hello", "hello"},
		{42, "42"},
		{3.14, "3.14"},
		{true, "true"},
		{false, "false"},
		{nil, "<nil>"},
	}

	for _, tt := range tests {
		result := Stringify(tt.input)
		if result != tt.expected {
			t.Errorf("Stringify(%v) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

// TestGetIntValue verifies GetIntValue function.
func TestGetIntValue(t *testing.T) {
	// Should work with int
	result := GetIntValue(42)
	if result != 42 {
		t.Errorf("GetIntValue(42) = %d, want 42", result)
	}

	// Should return 0 for non-int without panicking
	result = GetIntValue("not an int")
	if result != 0 {
		t.Errorf("GetIntValue(non-int) should return 0, got %d", result)
	}
}

// TestIsNumeric verifies IsNumeric function.
func TestIsNumeric(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
	}{
		{42, true},
		{3.14, true},
		{"hello", false},
		{true, false},
		{nil, false},
	}

	for _, tt := range tests {
		result := IsNumeric(tt.input)
		if result != tt.expected {
			t.Errorf("IsNumeric(%v) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

// TestSumInt verifies Sum with []int.
func TestSumInt(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	result := Sum(values)
	expected := 15.0

	if result != expected {
		t.Errorf("Sum([]int) = %.2f, want %.2f", result, expected)
	}
}

// TestSumFloat64 verifies Sum with []float64.
func TestSumFloat64(t *testing.T) {
	values := []float64{1.5, 2.5, 3.0}
	result := Sum(values)
	expected := 7.0

	if result != expected {
		t.Errorf("Sum([]float64) = %.2f, want %.2f", result, expected)
	}
}

// TestSumInterface verifies Sum with []interface{}.
func TestSumInterface(t *testing.T) {
	values := []interface{}{1, 2.5, 3}
	result := Sum(values)
	expected := 6.5

	if result != expected {
		t.Errorf("Sum([]interface{}) = %.2f, want %.2f", result, expected)
	}
}

// TestProcessValue verifies ProcessValue function.
func TestProcessValue(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{"hello", "String: hello"},
		{42, "Integer: 42"},
		{3.14, "Unknown"},
	}

	for _, tt := range tests {
		result := ProcessValue(tt.input)
		if result != tt.expected {
			t.Errorf("ProcessValue(%v) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

// TestGetTypeName verifies GetTypeName function.
func TestGetTypeName(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{42, "int"},
		{"hello", "string"},
		{3.14, "float64"},
		{true, "bool"},
		{[]int{1, 2, 3}, "[]int"},
	}

	for _, tt := range tests {
		result := GetTypeName(tt.input)
		if result != tt.expected {
			t.Errorf("GetTypeName(%v) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

// TestConvertToInt verifies ConvertToInt function.
func TestConvertToInt(t *testing.T) {
	tests := []struct {
		input       interface{}
		expected    int
		shouldError bool
	}{
		{42, 42, false},
		{3.14, 3, false},
		{"123", 123, false},
		{true, 1, false},
		{false, 0, false},
		{"not a number", 0, true},
		{[]int{}, 0, true},
	}

	for _, tt := range tests {
		result, err := ConvertToInt(tt.input)
		if tt.shouldError {
			if err == nil {
				t.Errorf("ConvertToInt(%v) should return error", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("ConvertToInt(%v) unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("ConvertToInt(%v) = %d, want %d", tt.input, result, tt.expected)
			}
		}
	}
}

// TestFilter verifies Filter function.
func TestFilter(t *testing.T) {
	values := []interface{}{1, "hello", 2, "world", 3.14, true, 3}

	// Filter ints
	ints := Filter(values, "int")
	if len(ints) != 3 {
		t.Errorf("Filter for ints: expected 3 values, got %d", len(ints))
	}

	// Filter strings
	strings := Filter(values, "string")
	if len(strings) != 2 {
		t.Errorf("Filter for strings: expected 2 values, got %d", len(strings))
	}

	// Filter floats
	floats := Filter(values, "float64")
	if len(floats) != 1 {
		t.Errorf("Filter for float64: expected 1 value, got %d", len(floats))
	}
}

// TestMax verifies Max function.
func TestMax(t *testing.T) {
	values := []interface{}{5, 2, 8, 1, 9, 3}
	result := Max(values)

	if result != 9 {
		t.Errorf("Max(%v) = %v, want 9", values, result)
	}
}

// TestMaxEmpty verifies Max with empty slice.
func TestMaxEmpty(t *testing.T) {
	values := []interface{}{}
	result := Max(values)

	if result != nil {
		t.Errorf("Max([]) should return nil, got %v", result)
	}
}

// TestTypeSafeMapString verifies TypeSafeMap with strings.
func TestTypeSafeMapString(t *testing.T) {
	m := NewTypeSafeMap()
	m.Set("name", "Alice")

	result := m.GetString("name")
	if result != "Alice" {
		t.Errorf("GetString() = %q, want %q", result, "Alice")
	}
}

// TestTypeSafeMapInt verifies TypeSafeMap with ints.
func TestTypeSafeMapInt(t *testing.T) {
	m := NewTypeSafeMap()
	m.Set("age", 30)

	result := m.GetInt("age")
	if result != 30 {
		t.Errorf("GetInt() = %d, want %d", result, 30)
	}
}

// TestTypeSafeMapContains verifies Contains method.
func TestTypeSafeMapContains(t *testing.T) {
	m := NewTypeSafeMap()
	m.Set("key1", "value1")
	m.Set("key2", nil)

	if !m.Contains("key1") {
		t.Error("Contains('key1') should return true")
	}

	if !m.Contains("key2") {
		t.Error("Contains('key2') should return true even if value is nil")
	}

	if m.Contains("key3") {
		t.Error("Contains('key3') should return false")
	}
}

// TestTypeSafeMapCount verifies Count method.
func TestTypeSafeMapCount(t *testing.T) {
	m := NewTypeSafeMap()
	m.Set("name", "Alice")
	m.Set("city", "NYC")
	m.Set("age", 30)
	m.Set("score", 95)

	stringCount := m.Count("string")
	if stringCount != 2 {
		t.Errorf("Count('string') = %d, want 2", stringCount)
	}

	intCount := m.Count("int")
	if intCount != 2 {
		t.Errorf("Count('int') = %d, want 2", intCount)
	}
}

// TestTypeSafeMapWrongType verifies handling of wrong type assertions.
func TestTypeSafeMapWrongType(t *testing.T) {
	m := NewTypeSafeMap()
	m.Set("value", 42)

	// Should not panic, should return empty string
	result := m.GetString("value")
	if result != "" {
		t.Errorf("GetString() on int value should return empty string, got %q", result)
	}
}
