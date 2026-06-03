func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	n := len(height)
	res := 0

	for i := 0; i < n; i++ {
		leftMax := height[i]
		rightMax := height[i]

		for j := 0; j < i; j++ {
			if height[j] > leftMax {
				leftMax = height[j]
			}
		}

		for j := i+1; j < n; j++ {
			if height[j] > rightMax {
				rightMax = height[j]
			}
		}

		res += min(leftMax, rightMax) - height[i]
	}

	return res
}
