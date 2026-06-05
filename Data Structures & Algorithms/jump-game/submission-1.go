func canJump(nums []int) bool {
    var memo = make(map[int]bool)
    var dfs func(i int) bool
    dfs = func (i int) bool {
        if res, ok := memo[i]; ok {
            return res
        }

        if (i == len(nums) - 1) {
            return true
        }

        if nums[i] == 0 {
            memo[i] = false
            return false
        }

        end := min(nums[i] + i, len(nums) - 1)
        for j := i + 1; j <= end; j++ {
            if dfs(j) {
                memo[i] = true
                return true
            }
        }

        memo[i] = false
        return false
    }

    return dfs(0)
}






















// recursion
    // var dfs func (i int) bool
    // dfs = func (i int) bool {
    //     if i == len(nums) - 1 {
    //         return true
    //     }
    //     end := min(len(nums) - 1, nums[i] + i)
    //     for j := i + 1; j <= end; j++ {
    //         if dfs(j) {
    //             return true
    //         }
    //     }

    //     return false
    // }

    // return dfs(0)