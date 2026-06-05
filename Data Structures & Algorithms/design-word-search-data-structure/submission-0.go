type WordDictionary struct {
	store []string
}

func Constructor() WordDictionary {
    return WordDictionary{store: []string{}}
}

func (this *WordDictionary) AddWord(word string)  {
    this.store = append(this.store, word)
}

func (this *WordDictionary) Search(word string) bool {
    for _, w := range this.store {
		if len(w) == len(word) {
			match := true
			for i := 0; i < len(w); i++ {
				if w[i] != word[i] && word[i] != '.' {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}

	return false
}
