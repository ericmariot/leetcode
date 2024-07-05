func maxProfit(prices []int) int {
    low := prices[0]
    maxProfit := 0

    for _, price := range prices {
        if price < low {
            low = price
        }

        if price-low > maxProfit {
            maxProfit = price-low
        }
    }


    return maxProfit
}
