func trap(height []int) int {
	n := len(height)
	res := 0
	if (n == 0) {
		return res
	}
	stack := []int{}
	
	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack) - 1]] {
			bottomIdx := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			if len(stack) == 0 {
				break
			}
			leftIdx := stack[len(stack) - 1]
			rightIdx := i

			distance := rightIdx - leftIdx - 1;
			boundedHeight := min(height[leftIdx], height[rightIdx]) - height[bottomIdx]

			res += distance * boundedHeight
		}

		stack = append(stack, i)
	}

	return res
}



























	// stack
	// n := len(height)
	// res := 0
	// if n == 0 {
	// 	return res
	// }
	// stack := make([]int, 0)
	// for i := 0; i < n; i++ {
	// 	for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
	// 		bottomIdx := stack[len(stack)-1]
	// 		stack = stack[:len(stack)-1]

	// 		if (len(stack) == 0) {
	// 			break
	// 		}
	// 		leftIdx := stack[len(stack)-1]
	// 		rightIdx := i
	// 		distance := i - leftIdx - 1
	// 		boundedHeight := min(height[rightIdx], height[leftIdx]) - height[bottomIdx]
	// 		res += distance * boundedHeight
	// 	}
	// 	stack = append(stack, i)
	// }

	// return res



	// two pointers
	// n := len(height)
	// res := 0
	// if n == 0 {
	// 	return res
	// } 
	// l, r := 0, n-1
	// leftMax, rightMax := height[l], height[r]
	// for l < r {
	// 	if leftMax < rightMax {
	// 		l++
	// 		leftMax = max(leftMax, height[l])
	// 		res += leftMax - height[l]
	// 	} else {
	// 		r--
	// 		rightMax = max(rightMax, height[r])
	// 		res += rightMax - height[r]
	// 	}
	// }

	// return res

	// prefix & suffix arrays
	// n := len(height)
	// if (n == 0) {
	// 	return 0
	// }
	// res := 0
	// leftMax := make([]int, n)
	// rightMax := make([]int, n)

	// leftMax[0] = height[0]
	// for i := 1; i < n; i++ {
	// 	leftMax[i] = max(leftMax[i-1], height[i])
	// }

	// rightMax[n-1] = height[n-1]
	// for i := n - 2; i >= 0; i-- {
	// 	rightMax[i] = max(rightMax[i+1], height[i])
	// }

	// for i := 0; i < n; i++ {
	// 	res += min(leftMax[i], rightMax[i]) - height[i]
	// }

	// return res

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