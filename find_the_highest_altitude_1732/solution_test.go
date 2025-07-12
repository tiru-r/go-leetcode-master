package find_the_highest_altitude_1732

import "testing"

func Test_largestAltitude(t *testing.T) {
	type args struct {
		gain []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample 1", args{gain: []int{-5, 1, 5, 0, -7}}, 1},
		{"sample 2", args{gain: []int{-4, -3, -2, -1, 4, 3, 2}}, 0},
		{"empty gain", args{gain: []int{}}, 0},
		{"only descents", args{gain: []int{-1, -2, -3}}, 0},
		{"only ascents", args{gain: []int{3, 2, 1}}, 6},
		{"single zero", args{gain: []int{0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestAltitude(tt.args.gain); got != tt.want {
				t.Errorf("largestAltitude() = %v, want %v", got, tt.want)
			}
		})
	}
}
