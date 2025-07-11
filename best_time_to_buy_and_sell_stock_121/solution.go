package best_time_to_buy_and_sell_stock_121

// O(n) time and O(1) space
func maxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	res := 0
	low := prices[0]
	for _, p := range prices {
		res = max(res, p-low)
		low = min(low, p)
	}
	return res
}

// O(n) time and O(1) space
func maxProfit2(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]
	// Modern range-over-int for cleaner, faster bounds checking
	for i := range len(prices) {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if maxProfit < prices[i]-minPrice {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

// O(n^2) time and O(1) space
func maxProfit(prices []int) int {
	maxProfit := 0
	// O(n²) approach with modern range syntax (still not recommended for production)
	for i := range len(prices) {
		buyPrice := prices[i]
		maxSellPrice := 0

		for j := range prices[i+1:] {
			sellPrice := prices[i+1+j]

			if sellPrice > buyPrice {
				maxSellPrice = max(maxSellPrice, sellPrice)
			}
		}

		maxProfit = max(maxProfit, maxSellPrice-buyPrice)
	}

	return maxProfit
}
