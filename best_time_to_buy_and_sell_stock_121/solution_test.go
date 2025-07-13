package best_time_to_buy_and_sell_stock_121

import "testing"

func TestMaxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Basic case", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
		{"Decreasing prices", args{[]int{7, 6, 4, 3, 1}}, 0},
		{"Single price", args{[]int{5}}, 0},
		{"Empty array", args{[]int{}}, 0},
		{"Two prices ascending", args{[]int{1, 5}}, 4},
		{"Two prices descending", args{[]int{5, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxProfit(tt.args.prices)
			if got != tt.want {
				t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
