func twoSum(nums []int, target int) []int {
    numsCnt := len(nums)
    numsOccs := make(map[int]int)
    res := []int{}

    for i := 0; i < numsCnt; i++ {
        currValue := nums[i]
        valToFind := target - currValue

        val, ok := numsOccs[valToFind]
        if ok {
            return []int{val, i}
        }

        val, ok = numsOccs[currValue]
        if !ok {
            numsOccs[currValue] = i
        }
    } 

    return res
}
