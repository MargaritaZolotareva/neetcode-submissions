func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int)
    var res []int

    for _, num := range nums {
        count[num]++
    }

    pairs := make([][2]int, 0, len(count))
    for num, cnt := range count {
        pairs = append(pairs, [2]int{cnt, num})
    }

    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i][0] > pairs[j][0]
    })

    for i := 0; i < k; i++ {
        res = append(res, pairs[i][1])
    }

    return res
}
