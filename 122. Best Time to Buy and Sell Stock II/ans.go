func maxProfit(prices []int) int {
    if len(prices) == 1 {
        return 0
    }
    prices = append(prices, -1)
    profit := 0
    min := prices[0]
    for i := 1; i < len(prices); i++ {
        if prices[i] < prices[i - 1] {
            profit += prices[i-1] - min
            min = prices[i]

        }
    }
    return profit
}
