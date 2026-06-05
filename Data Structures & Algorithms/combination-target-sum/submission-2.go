func combinationSum(nums []int, target int) [][]int {
    res := [][]int{}
    sort.Ints(nums)
    var dfs func(currSum int, combElems []int, i int)
    dfs = func(currSum int, combElems []int, i int) {
        if currSum == target {
            temp := make([]int, len(combElems))
            copy(temp, combElems)
            res = append(res, temp)
            return
        }

        for j := i; j < len(nums); j++ {
            if currSum + nums[j] > target {
                return
            }
            combElems = append(combElems, nums[j])
            dfs(currSum + nums[j], combElems, j)
            combElems = combElems[:len(combElems) - 1]
        }
    }

    dfs(0, []int{}, 0)

    return res
}
