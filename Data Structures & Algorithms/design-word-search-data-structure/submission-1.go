type TrieNode struct {
	children map[rune]*TrieNode
	isEnd bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: NewTrieNode()}
}

func (this *WordDictionary) AddWord(word string)  {
	cur := this.root
	for _, c := range word {
		if cur.children[c] == nil {
			cur.children[c] = NewTrieNode()
		}
		cur = cur.children[c]
	}
	cur.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.dfs(word, 0, this.root)
}

func (this *WordDictionary) dfs(word string, j int, root *TrieNode) bool {
	cur := root
	for i := j; i < len(word); i++ {
		c := rune(word[i])
		if c == '.' {
			for _, child := range cur.children {
				if child != nil && this.dfs(word, i + 1, child) {
					return true
				}
			}
			return false
		} else {
			if cur.children[c] == nil {
				return false
			}
			cur = cur.children[c]
		}
	}

	return cur.isEnd
}



















// type WordDictionary struct {
// 	store []string
// }

// func Constructor() WordDictionary {
//     return WordDictionary{store: []string{}}
// }

// func (this *WordDictionary) AddWord(word string)  {
//     this.store = append(this.store, word)
// }

// func (this *WordDictionary) Search(word string) bool {
//     for _, w := range this.store {
// 		if len(w) == len(word) {
// 			match := true
// 			for i := 0; i < len(w); i++ {
// 				if w[i] != word[i] && word[i] != '.' {
// 					match = false
// 					break
// 				}
// 			}
// 			if match {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }

