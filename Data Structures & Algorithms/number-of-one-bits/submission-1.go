func hammingWeight(n int) int {
	count := 0
	for i := 0; i < 32; i++ {
		if n&(1<<i) != 0 {
			count++
		}
	}
	return count
}







// bit mask
	// count := 0
	// for n != 0 {
	// 	n = n & (n-1)
	// 	count++
	// }

	// return count