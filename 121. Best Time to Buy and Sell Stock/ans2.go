func maxProfit(prices []int) int {
    if len(prices) == 0 || len(prices) == 1{
        return 0
    }

    profit := 0
    minBefore := prices[0]
    for i := 1; i < len(prices); i++ {
        profit = max(profit, prices[i] - minBefore)
        minBefore = min(minBefore, prices[i])
    }
    return profit
}
