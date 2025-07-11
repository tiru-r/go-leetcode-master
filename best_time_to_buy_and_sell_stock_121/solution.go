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
	for i := 0; i < len(prices); i++ {
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
	for i := 0; i < len(prices); i++ {
		buyPrice := prices[i]
		maxSellPrice := 0

		for j := i + 1; j < len(prices); j++ {
			sellPrice := prices[j]

			if sellPrice > buyPrice {
				maxSellPrice = max(maxSellPrice, sellPrice)
			}
		}

		maxProfit = max(maxProfit, maxSellPrice-buyPrice)
	}

	return maxProfit
}
