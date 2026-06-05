func maxArea(heights []int) int {
    hLen := len(heights)
    res := 0
    for i := 0; i < hLen; i++ {
        for j := i + 1; j < hLen; j++ {
            area := min(heights[i], heights[j]) * (j - i)
            res = max(area, res)
        }
    }

    return res
}





































    // maxSq := 0
    // left := 0
    // right := len(heights) - 1

    // for left < right {
    //     bottom := right - left
    //     side := min(heights[left], heights[right])
    //     area := bottom * side

    //     maxSq = max(maxSq, area)

    //     if (heights[left] < heights[right]) {
    //         left++
    //     } else {
    //         right--
    //     }
    // }

    // return maxSq