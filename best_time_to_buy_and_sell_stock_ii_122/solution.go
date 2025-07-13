package best_time_to_buy_and_sell_stock_ii_122

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	
	totalProfit := 0
	for i := range len(prices) - 1 {
		if diff := prices[i+1] - prices[i]; diff > 0 {
			totalProfit += diff
		}
	}
	return totalProfit
}
