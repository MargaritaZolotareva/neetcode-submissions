type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var res string
    for _, str := range strs {
        strLen := len(str)
        res += strconv.Itoa(strLen) + "#" + str
    }

    return res
}

func (s *Solution) Decode(encoded string) []string {
    res := []string{}
    cntInt := 0
    cnt := ""
    for i := 0; i < len(encoded); i++ {
        if (encoded[i] != '#') {
            cnt += string(encoded[i])
        } else {
            cntInt, _ = strconv.Atoi(cnt)
            cnt = ""
            endPos := i + cntInt + 1
            res = append(res, encoded[i + 1:endPos])
            i = endPos - 1
        }
    }

    return res
}
