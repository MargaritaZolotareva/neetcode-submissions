func threeSum(nums []int) [][]int {
    res := [][]int{}
    sort.Ints(nums)

    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            return res
        }

        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }

        l := i + 1
        r := len(nums) - 1

        for l < r {
            sum := nums[i] + nums[l] + nums[r]

            if sum == 0 {
                res = append(res, []int{nums[i], nums[l], nums[r]})
                l++
                r--
                for l < r && nums[l] == nums[l - 1] {
                    l++
                }
            } else {
                if sum < 0 {
                    l++
                } else {
                    r--
                }
            }
        }
    }
    
    return res
}
