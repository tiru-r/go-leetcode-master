package best_time_to_buy_and_sell_stock_ii_122

func maxProfit(prices []int) int {
	totalProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			totalProfit += prices[i] - prices[i-1] // Add profit if price increases
		}
	}
	return totalProfit
}
