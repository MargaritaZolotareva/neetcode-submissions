func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}
		mid := (right + left) / 2

		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
























	// brute force
	// minV := nums[0]

	// for _, val := range nums {
	// 	minV = min(minV, val)
	// }

	// return minV