// Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock
package besttimetobuyandsellstock

// Solution utilizing DP (top-down): O(n)

// func maxProfit(prices []int) int {
// 	n := len(prices)
// 	if n == 0 {
// 		return 0
// 	}
// 	memo := make([][2]int, n) // 2 indexes: not holding and holding stocks
// 	for i := range memo {
// 		memo[i][0] = -1 // -1 = not calc'ed yet
// 		memo[i][1] = -1
// 	}

// 	// DP function (recursive top-down)
// 	var dp func(i int, holding int) int
// 	dp = func(i int, holding int) int {
// 		if i == n {
// 			return 0
// 		}
// 		if memo[i][holding] != -1 {
// 			return memo[i][holding]
// 		}
// 		if holding == 1 {
// 			// Option 1: Keep holding the old stock
// 			hold := dp(i+1, 1)
// 			// Option 2: Sell the stock today with the current price
// 			sell := prices[i]
// 			memo[i][holding] = max(hold, sell)
// 		} else {
// 			// Option 1: Keep waiting without buying nothing
// 			wait := dp(i+1, 0)
// 			// Option 2: Buy the stock today
// 			buy := -prices[i] + dp(i+1, 1)
// 			memo[i][holding] = max(wait, buy)
// 		}
// 		return memo[i][holding]
// 	}

// 	return dp(0, 0)
// }

// Solution utilizing DP (bottom-up): O(n)

// func maxProfit(prices []int) int {
// 	n := len(prices)
// 	if n == 0 {
// 		return 0
// 	}

// 	dpBought := make([]int, n)
// 	dpSold := make([]int, n)

// 	dpBought[0] = -prices[0] // If I bought first day, I spent prices[0]
// 	dpSold[0] = 0            // If I don't bought anything, start with zero

// 	for i := 1; i < n; i++ {
// 		// To end the day with a long position:
// 		// Either I was already long yesterday (dpBought[i-1])
// 		// Or I decided to buy today from scratch (-prices[i])
// 		dpBought[i] = max(dpBought[i-1], -prices[i])

// 		// To end the day with a short position/no stock:
// 		// Either I already held no stock yesterday (dpVendido[i-1])
// 		// Or I sold the stock today that I had bought previously (dpBought[i-1] + prices[i])
// 		dpSold[i] = max(dpSold[i-1], dpBought[i-1]+prices[i])
// 	}
// 	return dpSold[n-1]
// }

// Solution utilizing Sliding Window: O(n)

func maxProfit(prices []int) int {
	best := 0
	for left, right := 0, 0; right < len(prices); right++ {
		if prices[left] > prices[right] {
			left = right
		} else {
			diff := prices[right] - prices[left]
			best = max(diff, best)
		}
	}
	return best
}
