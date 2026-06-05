func evalRPN(tokens []string) int {
	index := len(tokens) - 1

	var dfs func() int
	dfs = func() int {
		currSym := tokens[index]
		index--

		if currSym != "+" && currSym != "-" && currSym != "*" && currSym != "/" {
			num, _ := strconv.Atoi(currSym)
			return num
		}

		right := dfs()
		left := dfs()

		switch currSym {
			case "+":
				return left+right
			case "-":
				return left-right
			case "*":
				return left*right
			default:
				return left/right
		}
	}

	return dfs()
}























	// stack
	// stack := make([]int, 0)

	// for _, v := range tokens {
	// 	switch v {
	// 		case "+":
	// 			op1 := stack[len(stack) - 1]
	// 			op2 := stack[len(stack) - 2]
	// 			stack = stack[:len(stack) - 2]
	// 			stack = append(stack, op1+op2)
	// 		case "-":
	// 			op1 := stack[len(stack) - 1]
	// 			op2 := stack[len(stack) - 2]
	// 			stack = stack[:len(stack) - 2]
	// 			stack = append(stack, op2-op1)
	// 		case "/":
	// 			op1 := stack[len(stack) - 1]
	// 			op2 := stack[len(stack) - 2]
	// 			stack = stack[:len(stack) - 2]
	// 			stack = append(stack, op2/op1)
	// 		case "*":
	// 			op1 := stack[len(stack) - 1]
	// 			op2 := stack[len(stack) - 2]
	// 			stack = stack[:len(stack) - 2]
	// 			stack = append(stack, op2*op1)
	// 		default:
	// 			vi, _ := strconv.Atoi(v)
	// 			stack = append(stack, vi)
	// 	}
	// }

	// return stack[0]