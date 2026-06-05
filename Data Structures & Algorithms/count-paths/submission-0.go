func uniquePaths(m int, n int) int {
    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if i >= m || j >= n {
            return 0
        }

        if i == m - 1 && j == n - 1 {
            return 1
        }

        return dfs(i + 1, j) + dfs(i, j + 1)
    }

    return dfs(0, 0)
}