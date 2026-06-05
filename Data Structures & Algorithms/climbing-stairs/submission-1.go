func climbStairs(n int) int {
    memo := make([]int, n)
    for i := range memo {
        memo[i] = -1
    }
    var dfs func(i int) int
    dfs = func(i int) int {
        if i > n {
            return 0
        }
        if i == n {
            return 1
        }
        if memo[i] != -1 {
            return memo[i]
        }
        memo[i] = dfs(i + 1) + dfs(i + 2)
        return memo[i]
    }

    return dfs(0)
}
