func uniquePaths(m int, n int) int {
    var dfs func(i, j int) int
    memo := make([][]int, m)
    for i := range m {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
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