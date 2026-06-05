func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    chars := [26]int{}

    for k, v := range s {
        cs := v - 'a'
        chars[cs]++
        ct := t[k] - 'a'
        chars[ct]--
    }

    for _, v := range chars {
        if v != 0 {
            return false
        }
    }

    return true
}

// func isAnagram(s string, t string) bool {
//     occCntS := countOccurences(s)
//     occCntT := countOccurences(t)

//     return maps.Equal(occCntS, occCntT)
// }
// func countOccurences(s string) map[rune]int {
//     occCnt := make(map[rune]int)
    
//     for _, runeValue := range s {
//         _, ok := occCnt[runeValue]
//         if (ok) {
//             occCnt[runeValue]++
//         } else {
//             occCnt[runeValue] = 1
//         }
//     }

//     return occCnt
// }