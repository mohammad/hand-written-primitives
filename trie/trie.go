package trie

type Trie struct {
	Children map[byte]*Trie
	IsWord   bool
}

func Constructor() Trie {
	return Trie{
		Children: make(map[byte]*Trie),
		IsWord:   false,
	}
}

func (this *Trie) Insert(word string) {
	trie := this
	for i := 0; i < len(word); i++ {
		char := word[i]
		isWord := len(word)-1 == i

		val, ok := trie.Children[char]
		if ok {
			trie = val
		} else {
			// add it to the current trie
			trie.Children[char] = &Trie{
				Children: make(map[byte]*Trie),
			}
			trie = trie.Children[char]
		}
		if isWord == true {
			trie.IsWord = true
		}
	}
}

func (this *Trie) Search(word string) bool {
	trie := this
	for i := 0; i < len(word); i++ {
		char := word[i]
		val, ok := trie.Children[char]
		if !ok {
			return false
		}

		trie = val
	}

	return trie.IsWord
}

func (this *Trie) StartsWith(prefix string) bool {
	trie := this
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		val, ok := trie.Children[char]
		if !ok {
			return false
		}

		trie = val
	}

	return true
}
