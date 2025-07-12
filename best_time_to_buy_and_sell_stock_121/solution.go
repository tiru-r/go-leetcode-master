package best_time_to_buy_and_sell_stock_121

func maxProfit(prices []int) int {
	// Handle the edge case where the input array is empty.
	// If there are no prices, no profit can be made.
	if len(prices) == 0 {
		return 0
	}

	// Initialize maxProfit to 0. This will store the maximum profit found so far.
	maxProfit := 0

	// Initialize minPrice to the first element of the prices array.
	// This variable will keep track of the lowest price encountered up to the current day.
	minPrice := prices[0]

	// Iterate through each 'price' in the 'prices' array.
	for _, price := range prices {
		// If the current 'price' is less than 'minPrice', update 'minPrice'.
		// This means we've found a new, lower price to potentially buy at.
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			// If the current 'price' is not less than 'minPrice', it means we can potentially sell
			// at this 'price' given our current 'minPrice'.
			// Calculate the 'profit' if we were to buy at 'minPrice' and sell at the current 'price'.
			// If this calculated 'profit' is greater than the 'maxProfit' found so far, update 'maxProfit'.
			maxProfit = profit
		}
	}

	// After iterating through all prices, 'maxProfit' will hold the maximum possible profit.
	return maxProfit
}
