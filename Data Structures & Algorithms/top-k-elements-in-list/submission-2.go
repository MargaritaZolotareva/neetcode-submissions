func topKFrequent(nums []int, k int) []int {
    freqs := make(map[int]int)
    freqsBucket := make([][]int, len(nums) + 1)
    res := []int{}

    for _, val := range nums {
        freqs[val]++
    }

    for num, freq := range freqs {
        freqsBucket[freq] = append(freqsBucket[freq], num)
    }

    for i := len(freqsBucket) - 1; i >= 0; i-- {
        for _, elem := range freqsBucket[i] {
            res = append(res, elem)
        }
        if (len(res) == k) {
            return res
        }
    }

    return res
}


































// sorting
    // freqs := make(map[int]int)
    // var res []int

    // for _, num := range nums {
    //     freqs[num]++
    // }

    // pairs := make([][2]int, 0, len(freqs))
    // for num, freq := range freqs {
    //     pairs = append(pairs, [2]int{freq, num})
    // }
    
    // sort.Slice(pairs, func(i, j int) bool {
    //     return pairs[i][0] > pairs[j][0]
    // })

    // for i := 0; i < k; i++ {
    //     res = append(res, pairs[i][1])
    // }

    // return res