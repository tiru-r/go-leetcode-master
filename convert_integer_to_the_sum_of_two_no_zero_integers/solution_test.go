package convert_integer_to_the_sum_of_two_no_zero_integers

import (
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		wantA   int
		wantB   int
		wantErr bool
	}{
		{"n=2", 2, 1, 1, false},
		{"n=11", 11, 2, 9, false},
		{"n=100", 100, 1, 99, false},
		{"n=69", 69, 9, 60, false},
		{"n=1", 1, 0, 0, true},
		{"n=0", 0, 0, 0, true},
		{"n=-5", -5, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b, err := Split(tt.input)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Split(%d) err = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}

			if a+b != tt.input {
				t.Errorf("Split(%d) returned %d + %d = %d, want %d", tt.input, a, b, a+b, tt.input)
			}
			if hasZero(a) || hasZero(b) {
				t.Errorf("Split(%d) returned numbers containing '0': %d, %d", tt.input, a, b)
			}
		})
	}
}

func TestHasZero(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"123", 123, false},
		{"102", 102, true},
		{"0", 0, true},
		{"5", 5, false},
		{"-10", -10, true},
		{"100000", 100000, true},
		{"999999", 999999, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasZero(tt.input); got != tt.want {
				t.Errorf("hasZero(%d) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
