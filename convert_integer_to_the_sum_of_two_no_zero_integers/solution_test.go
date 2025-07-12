package convert_integer_to_the_sum_of_two_no_zero_integers

import (
	"reflect"
	"testing"
)

// TestGetNoZeroIntegers tests the GetNoZeroIntegers function for various inputs.
func TestGetNoZeroIntegers(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    []int
		wantErr bool
	}{
		{
			name:    "valid input n=2",
			input:   2,
			want:    []int{1, 1},
			wantErr: false,
		},
		{
			name:    "valid input n=11",
			input:   11,
			want:    []int{2, 9},
			wantErr: false,
		},
		{
			name:    "valid input n=100",
			input:   100,
			want:    []int{1, 99},
			wantErr: false,
		},
		{
			name:    "valid input n=69",
			input:   69,
			want:    []int{9, 60},
			wantErr: false,
		},
		{
			name:    "invalid input n=1",
			input:   1,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid input n=0",
			input:   0,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid input n=-5",
			input:   -5,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNoZeroIntegers(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNoZeroIntegers(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNoZeroIntegers(%d) = %v, want %v", tt.input, got, tt.want)
			}
			// Verify that the sum equals the input and neither number contains zero
			if err == nil {
				if got[0]+got[1] != tt.input {
					t.Errorf("GetNoZeroIntegers(%d) sum = %d, want %d", tt.input, got[0]+got[1], tt.input)
				}
				if hasZeroDigit(got[0]) || hasZeroDigit(got[1]) {
					t.Errorf("GetNoZeroIntegers(%d) returned numbers with zero digits: %v", tt.input, got)
				}
			}
		})
	}
}

// TestHasZeroDigit tests the hasZeroDigit function for various inputs.
func TestHasZeroDigit(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{
			name:  "no zero digit",
			input: 123,
			want:  false,
		},
		{
			name:  "has zero digit",
			input: 102,
			want:  true,
		},
		{
			name:  "zero input",
			input: 0,
			want:  true,
		},
		{
			name:  "single digit non-zero",
			input: 5,
			want:  false,
		},
		{
			name:  "negative number",
			input: -10,
			want:  true,
		},
		{
			name:  "large number with zero",
			input: 100000,
			want:  true,
		},
		{
			name:  "large number without zero",
			input: 999999,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasZeroDigit(tt.input)
			if got != tt.want {
				t.Errorf("hasZeroDigit(%d) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
