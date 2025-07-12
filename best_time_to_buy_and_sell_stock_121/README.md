# 121. Best Time to Buy and Sell Stock

https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

# Description

Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell 
one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:
```
Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
```

Example 2:
```
Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
```
# Explanation
-----

## Best Time to Buy and Sell Stock (LeetCode \#121)

This document explains the `maxProfit` function, which solves the "Best Time to Buy and Sell Stock" problem on LeetCode (problem \#121).

### Problem Description

The problem asks us to find the maximum profit that can be achieved by buying and selling a stock. We are given an array `prices` where `prices[i]` is the price of a given stock on the $i^{th}$ day. We are allowed to complete at most one transaction (i.e., buy one and sell one share of the stock). We cannot sell a stock before buying it.

### Function Explanation

The `maxProfit` function efficiently calculates the maximum profit by iterating through the `prices` array just once. It keeps track of the lowest price encountered so far and the maximum profit found.

```go
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

```

### How it Works (Step-by-Step)

1.  **Initialization**:

      * `maxProfit` is set to $0$. This is our initial assumption for the best profit.
      * `minPrice` is set to `prices[0]`. We assume the first day's price is the lowest buying opportunity initially.

2.  **Iteration**: The code then loops through each `price` in the `prices` array:

      * **Finding a Lower Buying Point**: If the current `price` is less than `minPrice`, it means we've found a new, better (lower) price to buy the stock. In this case, we update `minPrice` to this new `price`.
      * **Calculating Potential Profit**: If the current `price` is **not** less than `minPrice`, it means we might be able to make a profit by selling at the current `price`, having bought at our `minPrice`. We calculate `profit = price - minPrice`.
      * **Updating Maximum Profit**: If this `profit` is greater than the current `maxProfit`, it means we've found a new highest profit, so we update `maxProfit` accordingly.

3.  **Return Value**: After the loop finishes, `maxProfit` will contain the largest profit that could have been achieved by buying at the lowest point before a sale and selling at a higher point.

### Example

Let's consider `prices = [7, 1, 5, 3, 6, 4]`

| Day | Price | `minPrice` (after current price) | `profit` (`price - minPrice`) | `maxProfit` (after current price) | Explanation                                          |
| :-- | :---- | :------------------------------- | :---------------------------- | :-------------------------------- | :--------------------------------------------------- |
| 1   | 7     | 7                                | 0                             | 0                                 | Initialize `minPrice = 7`, `maxProfit = 0`         |
| 2   | 1     | 1                                | N/A                           | 0                                 | `price (1) < minPrice (7)`, update `minPrice = 1`    |
| 3   | 5     | 1                                | $5 - 1 = 4$                   | 4                                 | `price (5) > minPrice (1)`, `profit = 4`, `maxProfit = 4` |
| 4   | 3     | 1                                | $3 - 1 = 2$                   | 4                                 | `price (3) > minPrice (1)`, `profit = 2`, `maxProfit` remains 4 |
| 5   | 6     | 1                                | $6 - 1 = 5$                   | 5                                 | `price (6) > minPrice (1)`, `profit = 5`, `maxProfit = 5` |
| 6   | 4     | 1                                | $4 - 1 = 3$                   | 5                                 | `price (4) > minPrice (1)`, `profit = 3`, `maxProfit` remains 5 |

The function returns `5`.

### Time and Space Complexity

  * **Time Complexity**: $O(N)$, where $N$ is the number of prices. The function iterates through the `prices` array only once.
  * **Space Complexity**: $O(1)$. The function uses a constant amount of extra space for the `maxProfit` and `minPrice` variables.

-----