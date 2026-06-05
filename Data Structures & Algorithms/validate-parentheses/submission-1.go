func isValid(s string) bool {
    for strings.Contains(s, "()") || strings.Contains(s, "[]") || strings.Contains(s, "{}") {
        s = strings.ReplaceAll(s, "()", "")
        s = strings.ReplaceAll(s, "[]", "")
        s = strings.ReplaceAll(s, "{}", "")
    }

    return s == ""
}






























    // brackets := map[rune]rune{')': '(', ']': '[', '}': '{'}
    // stack := []rune{}

    // for _, val := range s {
    //     if openBracket, exists := brackets[val]; exists {
    //         if (len(stack) != 0) {
    //             elem := stack[len(stack)-1]
    //             stack = stack[:len(stack)-1]
    //             if elem != openBracket {
    //                 return false
    //             }
    //         } else {
    //             return false
    //         }
    //     } else {
    //         stack = append(stack, val)
    //     }
    // }

    // return len(stack) == 0