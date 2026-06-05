func isPalindrome(s string) bool {
    sLen := len(s)

    left := 0
    right := sLen - 1
    for left < right {
        for left < right && !isAlphaNum(rune(s[left])) {
            left++
        }
        for left < right && !isAlphaNum(rune(s[right])) {
            right--
        }

        if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
            return false
        }

        left++
        right--
    }

    return true
}

func isAlphaNum(c rune) bool {
    return unicode.IsLetter(c) || unicode.IsDigit(c)
}