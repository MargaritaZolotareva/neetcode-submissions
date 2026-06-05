func maxSubArray(nums []int) int {
    memo := make([][2]*int, len(nums) + 1)
    for i := range memo {
        memo[i] = [2]*int{nil, nil}
    }
    var dfs func (i int, flag int) int
    dfs = func (i int, flag int) int {
        if (i == len(nums)) {
            if flag == 1 {
                return 0
            } else {
                return -1e6
            }
        }
        if memo[i][flag] != nil {
            return *memo[i][flag]
        }

        if flag == 1 {
            res := max(0, nums[i] + dfs(i + 1, 1))
            memo[i][flag] = &res
            return res
        } else {
            res := max(dfs(i + 1, 0), nums[i] + dfs(i + 1, 1))
            memo[i][flag] = &res
            return res
        }
    }

    return dfs(0, 0)
}





























    // recursion
    // var dfs func(i int, flag bool) int

    // dfs = func(i int, flag bool) int {
    //     if (i == len(nums)) {
    //         if flag {
    //             return 0
    //         } else {
    //             return -1e6
    //         }
    //     }

    //     if flag {
    //         return max(0, nums[i] + dfs(i + 1, true))
    //     }

    //     return max(dfs(i + 1, false), dfs(i + 1, true) + nums[i])
    // }

    // return dfs(0, false)

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