class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        max_profit = 0
        min_price = prices[0]  # Start by assuming we buy on day 1
        
        for price in prices:
            # 1. Track the lowest buying price we've seen so far
            if price < min_price:
                min_price = price
            
            # 2. Check the profit if we sold today, and update max_profit if it's better
            current_profit = price - min_price
            if current_profit > max_profit:
                max_profit = current_profit
                
        return max_profit