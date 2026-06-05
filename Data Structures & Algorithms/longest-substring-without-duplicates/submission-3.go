func lengthOfLongestSubstring(s string) int {
    syms := make(map[byte]int)
    sLen := len(s)
    left := 0
    maxLen := 0

    for right := 0; right < sLen; right++ {
        if idx, found := syms[s[right]]; found {
            left = max(left, idx + 1)
        }
        maxLen = max(maxLen, right - left + 1)
        syms[s[right]] = right
    }

    return maxLen
}
