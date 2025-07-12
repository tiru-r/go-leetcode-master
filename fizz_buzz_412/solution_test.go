package fizz_buzz_412

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{"n = 0", 0, []string{}},
		{"n = 1", 1, []string{"1"}},
		{"n = 2", 2, []string{"1", "2"}},
		{"n = 3", 3, []string{"1", "2", "Fizz"}},
		{"n = 5", 5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{"n = 15", 15, []string{
			"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz",
			"Buzz", "11", "Fizz", "13", "14", "FizzBuzz",
		}},
		{"n = 16", 16, []string{
			"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz",
			"Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16",
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fizzBuzz(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzz(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
