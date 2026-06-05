func rob(nums []int) int {
    numsL := len(nums)
    memo := make([]int, numsL)
    for i := range memo {
        memo[i] = -1
    }
    var dfs func(i int) int
    dfs = func(i int) int {
        if (i >= numsL) {
            return 0
        }
        if memo[i] != -1 {
            return memo[i]
        }


        memo[i] = max(nums[i] + dfs(i+2), dfs(i+1))
        return memo[i]
    }
    return dfs(0)
}
