func countSubstrings(s string) int {
	manacher := func(s string) []int {
		l, r := 0, 0
		sLen := len(s)
		p := make([]int, sLen)
		for i:= 0; i < sLen; i++ {
			if i < r {
				p[i] = min(r - i, p[l + (r-i)])
			}

			for i + p[i] + 1 < sLen && i - p[i] - 1 >= 0 && s[i+p[i]+1] == s[i-p[i]-1] {
				p[i]++
			}

			if i+p[i] > r {
				l, r = i - p[i], i + p[i]
			}
		}

		return p
	} 

	res := 0
	sHash := "#" + joinWithSeparator(s, "#") + "#"
	p := manacher(sHash)

	for _, val := range p {
		res += (val + 1) / 2
	}

	return res
}

func joinWithSeparator(s, separator string) string {
	res := ""
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		res += string(s[i]) + separator
	}

	return res[:len(res) - 1]
}









	// manacher := func(s string) []int {
	// 	sHash := "#" + joinWithSeparator(s, "#") + "#"
	// 	sLen := len(sHash)
	// 	p := make([]int, sLen)
	// 	l, r := 0, 0
	// 	for i := 0; i < sLen; i++ {
	// 		if i < r {
	// 			p[i] = min(r - i, p[l+(r-i)])
	// 		}

	// 		for i + p[i] + 1 < sLen && i - p[i] - 1 >= 0 && sHash[i+p[i]+1] == sHash[i-p[i]-1] {
	// 			p[i]++
	// 		}

	// 		if i + p[i] > r {
	// 			l, r = i-p[i], i+p[i]
	// 		}
	// 	}

	// 	return p
	// }

	// p := manacher(s)
	// res := 0
	// for _, val := range p {
	// 	res += (val + 1) / 2
	// }

	// return res


	

// func joinWithSeparator(s, sep string) string {
// 	res := ""
// 	sLen := len(s)
// 	for i := 0; i < sLen; i++ {
// 		res += string(s[i]) + sep
// 	}

// 	return res[:len(res) - 1]
// }