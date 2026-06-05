import (
	"maps"
)

func isAnagram(s string, t string) bool {
    occCntS := countOccurences(s)
    occCntT := countOccurences(t)

    return maps.Equal(occCntS, occCntT)
}


func countOccurences(s string) map[rune]int {
    occCnt := make(map[rune]int)
    
    for _, runeValue := range s {
        _, ok := occCnt[runeValue]
        if (ok) {
            occCnt[runeValue]++
        } else {
            occCnt[runeValue] = 1
        }
    }

    return occCnt
}