func groupAnagrams(strs []string) [][]string {
    res := make(map[string][]string)
    strsCnt := len(strs)

    for i := 0; i < strsCnt; i++ {
        runes := []rune(strs[i])
        sort.Slice(runes, func(i,j int) bool {
            return runes[i] < runes[j]
        })
        
        sortedStr := string(runes)
        _, ok := res[sortedStr]
        if ok {
            res[sortedStr] = append(res[sortedStr], strs[i])
        } else {
            res[sortedStr] = []string{strs[i]}
        }
    }
	var result [][]string
	for _, group := range res {
		result = append(result, group)
	}

	return result
}
