func longestCommonSubsequence(text1 string, text2 string) int {
    m := len(text1)
    n := len(text2)
    dp := make([][]int, m + 1)
    for i := range dp {
        dp[i] = make([]int, n + 1)
    }

    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if (text1[i] == text2[j]) {
                dp[i][j] = 1 + dp[i+1][j+1]
            } else {
                dp[i][j] = max(dp[i+1][j], dp[i][j+1])
            }
        }
    }

    return dp[0][0]
}


























// DP bottom-up memo
    // m, n := len(text1), len(text2)
    // var dfs func(i, j int) int
    // memo := make([][]int, m)
    // for i := range memo {
    //     memo[i] = make([]int, n)
    //     for j := range memo[i] {
    //         memo[i][j] = -1
    //     }
    // }

    // dfs = func(i, j int) int {
    //     if i >= m || j >= n {
    //         return 0
    //     }
    //     if memo[i][j] != -1 {
    //         return memo[i][j]
    //     }

    //     if text1[i] == text2[j] {
    //         memo[i][j] = 1 + dfs(i + 1, j + 1)
    //     } else {
    //         memo[i][j] = max(dfs(i + 1, j), dfs(i, j + 1))
    //     }

    //     return memo[i][j]
    // }

    // return dfs(0, 0)