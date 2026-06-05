func uniquePaths(m int, n int) int {
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if i >= m || j >= n {
            return 0
        }

        if i == m - 1 && j == n - 1 {
            return 1
        }
        if memo[i][j] != -1 {
            return memo[i][j]
        }

        memo[i][j] = dfs(i + 1, j) + dfs(i, j + 1)
        
        return memo[i][j]
    }

    return dfs(0, 0)
}























    // DP Bottom-Up
    // arr := make([][]int, m + 1)
    // for i := range arr {
    //     arr[i] = make([]int, n + 1)
    // }
    // arr[m-1][n-1] = 1
    // for i := m - 1; i >= 0; i-- {
    //     for j := n - 1; j >= 0; j-- {
    //         arr[i][j] += arr[i + 1][j] + arr[i][j + 1]
    //     }
    // }

    // return arr[0][0]


// top-down (memoization)
    // var dfs func(i, j int) int
    // memo := make([][]int, m)
    // for i := range m {
    //     memo[i] = make([]int, n)
    //     for j := range memo[i] {
    //         memo[i][j] = -1
    //     }
    // }
    // dfs = func(i, j int) int {
    //     if i >= m || j >= n {
    //         return 0
    //     }

    //     if i == m - 1 && j == n - 1 {
    //         return 1
    //     }

    //     if memo[i][j] != -1 {
    //         return memo[i][j]
    //     }
    //     memo[i][j] = dfs(i + 1, j) + dfs(i, j + 1)
    //     return memo[i][j]
    // }

    // return dfs(0, 0)