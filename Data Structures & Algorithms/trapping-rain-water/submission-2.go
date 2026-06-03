func trap(height []int) int {
	// prefix & suffix arrays
	n := len(height)
	if (n == 0) {
		return 0
	}
	res := 0
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i := 0; i < n; i++ {
		res += min(leftMax[i], rightMax[i]) - height[i]
	}

	return res
}

























// brute force

	// if len(height) == 0 {
	// 	return 0
	// }

	// n := len(height)
	// res := 0

	// for i := 0; i < n; i++ {
	// 	leftMax := height[i]
	// 	rightMax := height[i]

	// 	for j := 0; j < i; j++ {
	// 		if height[j] > leftMax {
	// 			leftMax = height[j]
	// 		}
	// 	}

	// 	for j := i+1; j < n; j++ {
	// 		if height[j] > rightMax {
	// 			rightMax = height[j]
	// 		}
	// 	}

	// 	res += min(leftMax, rightMax) - height[i]
	// }

	// return res