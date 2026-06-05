func lengthOfLIS(nums []int) int {
    numsL := len(nums)
    var dfs func(i, j int) int

    dfs = func(i, j int) int {
        if i >= numsL {
            return 0
        }
        if j == -1 || nums[j] < nums[i] {
            return max(dfs(i + 1, j), 1 + dfs(i + 1, i))
        }

        return dfs(i + 1, j)
    }

    return dfs(0, -1)
}
