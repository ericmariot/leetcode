class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        cheaper = prices[0]
        profit = 0

        for price in prices:
            if price < cheaper:
                cheaper = price
            if price - cheaper > profit:
                profit = price - cheaper

        return profit
