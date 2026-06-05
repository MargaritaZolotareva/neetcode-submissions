func maxSubArray(nums []int) int {
    res, currSum := nums[0], 0

    for _, val := range nums {
        if currSum < 0 {
            currSum = 0
        }
        
        currSum += val
        res = max(currSum, res)
    }

    return res
}
