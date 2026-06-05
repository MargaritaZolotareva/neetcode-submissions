func findMin(nums []int) int {
	minV := nums[0]

	for _, val := range nums {
		minV = min(minV, val)
	}

	return minV
}
