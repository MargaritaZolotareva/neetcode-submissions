func longestPalindrome(s string) string {
	sLen := len(s)
	resIdx, resLen := 0, 0

	dp := make([][]bool, sLen)
	for i := range dp {
		dp[i] = make([]bool, sLen)
	}
	for i :=  sLen - 1; i >= 0; i-- {
		for j := i; j < sLen; j++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if resLen < (j - i + 1) {
					resLen = j - i + 1
					resIdx = i
				}
			}
		}
	}

	return s[resIdx : resIdx+resLen]
}

















	// brute force
    // result := ""
	// resLen := 0
	// sLen := len(s)

	// for i := 0; i < sLen; i++ {
	// 	for j := i; j < sLen; j++ {
	// 		l, r := i, j
	// 		for l < r && s[l] == s[r] {
	// 			l++
	// 			r--
	// 		}

	// 		if l >= r && resLen < (j - i + 1) {
	// 			 result = s[i : j+1]
	// 			 resLen = j - i + 1

	// 		}
	// 	}
	// }

	// return result