func canJump(nums []int) bool {
    var dfs func (i int) bool
    dfs = func (i int) bool {
        if i == len(nums) - 1 {
            return true
        }
        end := min(len(nums) - 1, nums[i] + i)
        for j := i + 1; j <= end; j++ {
            if dfs(j) {
                return true
            }
        }

        return false
    }

    return dfs(0)
}
