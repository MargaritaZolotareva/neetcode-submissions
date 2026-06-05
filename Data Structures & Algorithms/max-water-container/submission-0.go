func maxArea(heights []int) int {
    maxSq := 0
    left := 0
    right := len(heights) - 1

    for left < right {
        bottom := right - left
        side := min(heights[left], heights[right])
        area := bottom * side

        maxSq = max(maxSq, area)

        if (heights[left] < heights[right]) {
            left++
        } else {
            right--
        }
    }

    return maxSq
}
