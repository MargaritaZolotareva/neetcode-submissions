func maxProfit(prices []int) int {
// DP
    res := 0
    minEl := 101
    for _, price := range prices {
        currP := price - minEl
        res = max(res, currP)
        if (minEl > price) {
            minEl = price
        }
    }

    return res
}





















// two pointers
    // pricesLen := len(prices)
    // left := 0
    // right := 1
    // res := 0

    // for right < pricesLen {
    //     if (prices[left] < prices[right]) {
    //         currP := prices[right] - prices[left]
    //         res = max(res, currP)
    //     } else {
    //         left = right
    //     }
    //     right++
    // }

    // return res