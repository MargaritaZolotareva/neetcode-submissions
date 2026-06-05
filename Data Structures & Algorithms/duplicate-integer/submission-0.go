func hasDuplicate(nums []int) bool {
    numsCnt := len(nums)
    freqCnt := make(map[int]int)

    for i := 0; i < numsCnt; i++ {
        _, ok := freqCnt[nums[i]]
        if ok {
            return true
        } else {
            freqCnt[nums[i]] = 1
        }
    }

    return false
}
