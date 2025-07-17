package valid_number_65

import "testing"

func TestIsNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
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
		{"spaces", " 1 ", true},

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
		{"invalid_chars", "1#2", false},
		{"leading_dot_no_digits", "-.", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNumber(tt.input)
			if got != tt.expected {
				t.Errorf("isNumber(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
