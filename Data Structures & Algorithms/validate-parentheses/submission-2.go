type Stack struct {
    data []rune
}

func (s *Stack) Push(c rune) {
    s.data = append(s.data, c)
}

func (s *Stack) Pop() (rune, error) {
    item := s.data[len(s.data) - 1]
    s.data = s.data[:len(s.data) - 1]

    return item, nil
}

func (s *Stack) Empty() bool {
    return len(s.data) == 0
}

func isValid(s string) bool {
    stack := Stack{}
    closingBraces := map[rune]rune{')': '(', ']': '[', '}': '{'}

    for _, c := range s {
        if _, exists := closingBraces[c]; exists {
            if !stack.Empty() {
                brace, _ := stack.Pop()
                if brace != rune(closingBraces[c]) {
                    return false
                }
            } else {
                return false
            }
        } else {
            stack.Push(c)
        }
    }

    return stack.Empty()
}






























    // brackets := map[rune]rune{')': '(', ']': '[', '}': '{'}
    // stack := []rune{}

    // for _, val := range s {
    //     if openBracket, exists := brackets[val]; exists {
    //         if (len(stack) != 0) {
    //             elem := stack[len(stack)-1]
    //             stack = stack[:len(stack)-1]
    //             if elem != openBracket {
    //                 return false
    //             }
    //         } else {
    //             return false
    //         }
    //     } else {
    //         stack = append(stack, val)
    //     }
    // }

    // return len(stack) == 0