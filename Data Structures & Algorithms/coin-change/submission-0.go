func coinChange(coins []int, amount int) int {
    var dfs func(sum int) int
    memo := make(map[int]int)
    memo[0] = 0
    dfs = func(sum int) int {
        if val, ok := memo[sum]; ok {
            return val
        }
        res := math.MaxInt32
        for _, coin := range coins {
            if sum - coin >= 0 {
                curr := 1 + dfs(sum - coin)
                res = min(res, curr)
            }
        }
        memo[sum] = res

        return res
    }

    res := dfs(amount)
    if res >= math.MaxInt32 {
        return -1
    }

    return res
}
