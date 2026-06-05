func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	res := nums[0]

	for left <= right {
		if nums[left] < nums[right] {
			if nums[left] < res {
				return nums[left]
			}
		}

		mid := (left + right) / 2
		res = min(res, nums[mid])

		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	} 
	return res
}






























	// Binary Search (Lower Bound)
	// left := 0
	// right := len(nums) - 1

	// for left < right {
	// 	if nums[left] < nums[right] {
	// 		return nums[left]
	// 	}
	// 	mid := (right + left) / 2

	// 	if nums[mid] >= nums[left] {
	// 		left = mid + 1
	// 	} else {
	// 		right = mid
	// 	}
	// }

	// return nums[left]


	// Brute force
	// minV := nums[0]

	// for _, val := range nums {
	// 	minV = min(minV, val)
	// }

	// return minV