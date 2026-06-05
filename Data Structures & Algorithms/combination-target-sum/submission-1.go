func combinationSum(nums []int, target int) [][]int {
    res := [][]int{}
    var dfs func(currSum int, combElems []int, i int)
    dfs = func(currSum int, combElems []int, i int) {
        if currSum == target {
            temp := make([]int, len(combElems))
            copy(temp, combElems)
            res = append(res, temp)
            return
        }

        if i >= len(nums) || currSum > target {
            return
        }

        combElems = append(combElems, nums[i])
        dfs(currSum + nums[i], combElems, i)
        combElems = combElems[:len(combElems) - 1]
        dfs(currSum, combElems, i + 1)
    }

    dfs(0, []int{}, 0)

    return res
}
