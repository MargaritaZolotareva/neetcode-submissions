func maxSubArray(nums []int) int {
    var dfs func(i int, flag bool) int

    dfs = func(i int, flag bool) int {
        if (i == len(nums)) {
            if flag {
                return 0
            } else {
                return -1e6
            }
        }

        if flag {
            return max(0, nums[i] + dfs(i + 1, true))
        }

        return max(dfs(i + 1, false), dfs(i + 1, true) + nums[i])
    }

    return dfs(0, false)
}





























    //Kadane
    // res, currSum := nums[0], 0

    // for _, val := range nums {
    //     if currSum < 0 {
    //         currSum = 0
    //     }
        
    //     currSum += val
    //     res = max(currSum, res)
    // }

    // return res