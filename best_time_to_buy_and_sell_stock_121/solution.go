package best_time_to_buy_and_sell_stock_121

// MaxProfit finds maximum profit from single buy-sell transaction
// Time: O(n), Space: O(1)
func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices[1:] {
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
