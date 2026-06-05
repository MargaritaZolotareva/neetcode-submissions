type TrieNode struct {
	children map[byte]*TrieNode
	isWord bool
}

func NewTrieNode() *TrieNode{
	return &TrieNode{children: make(map[byte]*TrieNode)}
}

func (this *TrieNode) AddWord(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		c := word[i]
		if cur.children[c] == nil {
			cur.children[c] = NewTrieNode()
		}
		cur = cur.children[c]
	}

	cur.isWord = true
}

func findWords(board [][]byte, words []string) []string {
    root := NewTrieNode()
	for _, w := range words {
		root.AddWord(w)
	}

	rows, cols := len(board), len(board[0])
	var res []string
	visited := make(map[[2]int]bool)
	wordSet := make(map[string]bool)

	var dfs func(r, c int, node *TrieNode, word string)
	dfs = func(r, c int, node *TrieNode, word string) {
		if r < 0 || c < 0 || r >= rows || c >= cols || visited[[2]int{r,c}] || board[r][c] == 0 {
			return
		}

		char := board[r][c]
		nextNode, found := node.children[char]
		if !found {
			return
		}

		visited[[2]int{r, c}] = true
		word += string(char)
		if nextNode.isWord {
			wordSet[word] = true
		}
		dfs(r+1, c, nextNode, word);
		dfs(r-1, c, nextNode, word);
		dfs(r, c+1, nextNode, word);
		dfs(r, c-1, nextNode, word);

		visited[[2]int{r, c}] = false
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dfs(r, c, root, "")
		}
	}

	for word := range wordSet {
		res = append(res, word)
	}

	return res
}
