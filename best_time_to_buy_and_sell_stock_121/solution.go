package best_time_to_buy_and_sell_stock_121

// MaxProfit finds maximum profit from single buy-sell transaction
// Time: O(n), Space: O(1)
func MaxProfit(prices []int) int {
	// Return 0 if we have 0 or 1 price (can't make any transaction)
	if len(prices) <= 1 {
		return 0
	}

	// Initialize minimum price seen so far to first price
	minPrice := prices[0]
	// Initialize maximum profit to 0 (worst case: no profit)
	maxProfit := 0

	// Iterate through prices starting from second day
	for _, price := range prices[1:] {
		// If current price is lower than minimum seen, update minimum
		if price < minPrice {
			minPrice = price  // New lowest price to buy at
		} else if profit := price - minPrice; profit > maxProfit {
			// Calculate profit if we sell at current price
			// If this profit is better than our best so far, update it
			maxProfit = profit
		}
	}

	// Return the maximum profit achievable
	return maxProfit
}
