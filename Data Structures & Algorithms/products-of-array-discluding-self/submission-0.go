func productExceptSelf(nums []int) []int {
    n := len(nums)
    output := make([]int, n)
    currProd := 1
    output[0] = 1
    for i := 1; i < n; i++ {
        currProd *= nums[i - 1]
        output[i] = currProd
    }
    

    currProd = 1
    for i := n - 2; i >= 0; i-- {
        currProd *= nums[i + 1]
        output[i] *= currProd
    }

    return output
}
