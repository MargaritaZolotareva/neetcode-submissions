func rob(nums []int) int {
    numsL := len(nums)
    if numsL == 1 {
        return nums[0]
    }
    memo := make([][2]int, numsL)
    for i := range memo {
        memo[i][0] = -1
        memo[i][1] = -1
    }
    var dfs func(i, firstPicked int) int
    dfs = func(i, firstPicked int) int {
        if i >= numsL || firstPicked == 1 && i == numsL - 1 {
            return 0
        }
        if memo[i][firstPicked] != -1 {
            return memo[i][firstPicked]
        }
        // iZero := firstPicked
        // if (i == 0) {
        //     iZero = 1
        // }
        memo[i][firstPicked] = max(nums[i] + dfs(i + 2, firstPicked), dfs(i + 1, firstPicked))

        return memo[i][firstPicked]
    }

    return max(dfs(0, 1), dfs(1, 0))
}
