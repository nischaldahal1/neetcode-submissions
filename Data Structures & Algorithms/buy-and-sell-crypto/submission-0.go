func maxProfit(prices []int) int {
	max_profit := 0
	min_seen := prices[0]
	for _,price := range prices {
		min_seen = min(min_seen, price)
		max_profit = max(max_profit, price - min_seen)
	}
	return max_profit
}
