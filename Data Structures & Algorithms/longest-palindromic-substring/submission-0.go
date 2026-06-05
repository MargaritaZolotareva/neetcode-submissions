func longestPalindrome(s string) string {
    result := ""
	resLen := 0
	sLen := len(s)

	for i := 0; i < sLen; i++ {
		for j := i; j < sLen; j++ {
			l, r := i, j
			for l < r && s[l] == s[r] {
				l++
				r--
			}

			if l >= r && resLen < (j - i + 1) {
				 result = s[i : j+1]
				 resLen = j - i + 1

			}
		}
	}

	return result
}
