func longestPalindrome(s string) string {
	sLen := len(s)
	dp := make([][]bool, sLen)
	for i := range dp {
		dp[i] = make([]bool, sLen)
	}
	resIdx, resLen := 0, 0

	for i := 0; i < sLen; i++ {
		for j := i; j >= 0; j-- {
			if s[i] == s[j] && (i - j <= 2 || dp[i-1][j+1]) {
					dp[i][j] = true
					if resLen < (i-j+1) {
						resLen = i-j+1
						resIdx = j
					}
			}
		}
	}

	return s[resIdx:resIdx+resLen]
}























	// sLen := len(s)
	// resLen := 0
	// resInd := 0
	// dp := make([][]bool, sLen)
	// for i := range dp {
	// 	dp[i] = make([]bool, sLen)
	// }

	// // for i := sLen - 1; i >= 0; i-- {
	// // 	for j := i; j < sLen; j++ {
	// // 		if s[i] == s[j] && (j - i <= 2 || dp[i+1][j-1]) {
	// // 			dp[i][j] = true
	// // 			if j - i + 1 > resLen {
	// // 				resLen = j - i + 1
	// // 				resInd = i
	// // 			}
	// // 		}
	// // 	}
	// // }
	// for i := 0; i < sLen; i++ {
	// 	for j := i; j >= 0; j-- {
	// 		if s[i] == s[j] && (i - j <= 2 || dp[i-1][j+1]) {
	// 			dp[i][j] = true
	// 			if i - j + 1 > resLen {
	// 				resLen = i - j + 1
	// 				resInd = j
	// 			}
	// 		}
	// 	}
	// }

	// return s[resInd:resInd+resLen]


// two pointers
	// sLen := len(s)
	// resIdx, resLen := 0, 0

	
	// for i := 0; i < sLen; i++ {
	// 	l, r := i, i
	// 	for l >= 0 && r < sLen && s[l] == s[r] {
	// 		if resLen < r - l + 1 {
	// 			resIdx = l
	// 			resLen = r - l + 1
	// 		}
	// 		l--
	// 		r++
	// 	}

	// 	l, r = i, i+1
	// 	for l >= 0 && r < sLen && s[l] == s[r] {
	// 		if r - l + 1 > resLen {
	// 			resLen = r - l + 1
	// 			resIdx = l
	// 		}
	// 		l--
	// 		r++
	// 	}
	// }

	// return s[resIdx:resIdx+resLen]

//DP
	// sLen := len(s)
	// resIdx, resLen := 0, 0

	// dp := make([][]bool, sLen)
	// for i := range dp {
	// 	dp[i] = make([]bool, sLen)
	// }
	// for i :=  sLen - 1; i >= 0; i-- {
	// 	for j := i; j < sLen; j++ {
	// 		if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
	// 			dp[i][j] = true
	// 			if resLen < (j - i + 1) {
	// 				resLen = j - i + 1
	// 				resIdx = i
	// 			}
	// 		}
	// 	}
	// }

	// return s[resIdx : resIdx+resLen]

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