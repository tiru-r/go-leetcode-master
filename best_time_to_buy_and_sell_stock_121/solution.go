package best_time_to_buy_and_sell_stock_121

// maxProfit calculates the maximum profit that can be achieved by buying and selling
// a stock once, given an array of daily stock prices. The buy must occur before the sell.
// If no profit is possible, it returns 0.
func maxProfit(prices []int) int {
	// Handle edge cases where the input array has 0 or 1 element.
	// No profit can be made in these cases.
	if len(prices) <= 1 {
		return 0
	}

	// Initialize maxProfit to 0 to store the maximum profit found so far.
	maxProfit := 0

	// Initialize minPrice to the first element of the prices array.
	// This tracks the lowest price encountered up to the current day.
	minPrice := prices[0]

	// Iterate through each price in the prices array.
	for _, price := range prices {
		// If the current price is less than minPrice, update minPrice.
		// This represents a new potential buy point.
		if price < minPrice {
			minPrice = price
		} else if currentProfit := price - minPrice; currentProfit > maxProfit {
			// Calculate the profit if we buy at minPrice and sell at the current price.
			// Update maxProfit if this profit is greater than the current maxProfit.
			maxProfit = currentProfit
		}
	}

	// Return the maximum profit found.
	return maxProfit
}
