func canJump(nums []int) bool {
    numsLen := len(nums)
    goal := numsLen - 1

    for i := numsLen - 2; i >= 0; i-- {
        if (nums[i] + i >= goal) {
            goal = i
        }
        if (goal <= 0) {
            return true
        }
    }

    return goal == 0
}








































    // greedy
    // numsLen := len(nums)
    // goal := numsLen - 1

    // for i := numsLen - 2; i >= 0; i-- {
    //     if i + nums[i] >= goal {
    //         goal = i
    //     }
    // }

    // return goal == 0 
// DP Bottom-Up
    // n := len(nums)
    // dp := make([]bool, n)
    // dp[n - 1] = true

    // for i := n - 2; i >= 0; i-- {
    //     end := min(nums[i] + i, n)

    //     for j := i + 1 ; j <= end; j++ {
    //         if dp[j] {
    //             dp[i] = true
    //             break
    //         }
    //     }
    // }

    // return dp[0]

// DP Top-Down (recursion with memoization)
    // var memo = make(map[int]bool)
    // var dfs func(i int) bool
    // dfs = func (i int) bool {
    //     if res, ok := memo[i]; ok {
    //         return res
    //     }

    //     if (i == len(nums) - 1) {
    //         return true
    //     }

    //     end := min(nums[i] + i, len(nums) - 1)
    //     for j := i + 1; j <= end; j++ {
    //         if dfs(j) {
    //             memo[i] = true
    //             return true
    //         }
    //     }

    //     memo[i] = false
    //     return false
    // }

    // return dfs(0)

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