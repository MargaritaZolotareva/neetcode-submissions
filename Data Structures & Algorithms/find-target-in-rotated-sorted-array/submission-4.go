func search(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	pivot := 0

	for left < right {
		mid := (right+left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	pivot = left

	var binarySearch func (left, right int) int
	binarySearch = func(left, right int) int {
		for left <= right {
			mid := (left+right)/2
			if target == nums[mid] {
				return mid
			} else if target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		return -1
	}
	res := binarySearch(0, pivot - 1)
	if res != -1 {
		return res
	}
	return binarySearch(pivot, len(nums) - 1)
}
