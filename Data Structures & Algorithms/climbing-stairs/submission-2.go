func climbStairs(n int) int {
    // DP bottom-up
    if n < 3 {
        return n
    }
    dp := make([]int, n + 1)
    dp[n], dp[n - 1] = 1, 1

    for i := n - 2; i >= 0; i-- {
        dp[i] = dp[i+1] + dp[i+2]
    }

    return dp[0]
}

















    // DP top-down
    // memo := make([]int, n)
    // for i := range memo {
    //     memo[i] = -1
    // }
    // var dfs func(i int) int
    // dfs = func(i int) int {
    //     if i > n {
    //         return 0
    //     }
    //     if i == n {
    //         return 1
    //     }
    //     if memo[i] != -1 {
    //         return memo[i]
    //     }
    //     memo[i] = dfs(i + 1) + dfs(i + 2)
    //     return memo[i]
    // }

    // return dfs(0)