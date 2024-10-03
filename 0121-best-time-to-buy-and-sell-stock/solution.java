class Solution {
    public int maxProfit(int[] prices) {
        int cheaper = prices[0];
        int profit = 0;
        
        for (int i : prices) {
            if (i < cheaper) {
                cheaper = i;
            }

            if (i  - cheaper > profit) {
                profit = i - cheaper;
            }
        }

        return profit;
    }
}
