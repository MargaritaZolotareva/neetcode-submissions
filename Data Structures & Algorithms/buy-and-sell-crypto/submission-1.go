func maxProfit(prices []int) int {
    pricesLen := len(prices)
    left := 0
    right := 1
    res := 0

    for right < pricesLen {
        if (prices[left] < prices[right]) {
            currP := prices[right] - prices[left]
            res = max(res, currP)
        } else {
            left = right
        }
        right++
    }

    return res
}
