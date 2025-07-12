package valid_number_65

import (
	"fmt"
	"testing"
)

// TestIsNumber tests the IsNumber function with a variety of valid and invalid inputs
// as defined by LeetCode 65 "Valid Number".
func TestIsNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Valid numbers
		{"simple_integer", "2", true},
		{"zero_padded_integer", "0089", true},
		{"negative_decimal", "-0.1", true},
		{"positive_decimal", "+3.14", true},
		{"integer_with_trailing_dot", "4.", true},
		{"leading_dot_decimal", "-.9", true},
		{"exponent_positive", "2e10", true},
		{"exponent_negative", "-90E3", true},
		{"exponent_with_sign", "3e+7", true},
		{"exponent_decimal", "6e-1", true},
		{"single_zero", "0", true},
		{"decimal_with_zeros", "0.0", true},
		{"negative_zero", "-0", true},
		{"positive_zero", "+0", true},

		// Invalid numbers
		{"empty_string", "", false},
		{"non_numeric", "abc", false},
		{"number_with_letter", "1a", false},
		{"incomplete_exponent", "1e", false},
		{"exponent_without_number", "e3", false},
		{"decimal_exponent", "99e2.5", false},
		{"double_sign", "--6", false},
		{"conflicting_signs", "-+3", false},
		{"mixed_chars", "95a54e53", false},
		{"single_dot", ".", false},
		{"single_sign", "+", false},
		{"double_dot", "1.2.3", false},
		{"exponent_without_digits", "1e+", false},
		{"multiple_e", "1e2e3", false},
		{"spaces", " 1 ", false},
		{"invalid_chars", "1#2", false},
		{"leading_dot_no_digits", "-.", false},
		{"trailing_e", "1.", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsNumber([]byte(tt.input))
			if got != tt.expected {
				t.Errorf("IsNumber(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

// TestIsNumberNilInput tests the behavior of IsNumber with a nil input.
func TestIsNumberNilInput(t *testing.T) {
	if got := IsNumber(nil); got != false {
		t.Errorf("IsNumber(nil) = %v, want false", got)
	}
}

// ExampleIsNumber demonstrates usage of the IsNumber function.
func ExampleIsNumber() {
	fmt.Println(IsNumber([]byte("2.0")))   // true
	fmt.Println(IsNumber([]byte("-1e10"))) // true
	fmt.Println(IsNumber([]byte("abc")))   // false
	// Output:
	// true
	// true
	// false
}
