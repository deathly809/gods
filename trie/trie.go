
package trie

// Trie exports the common interface all tries have
type Trie interface {
	Insert(string)
	Remove(string)
	Contains(string)
	Count()
}

type trie struct {
	word bool
	children std::map[rune]*trie
}

func (t *trie) Insert(str string) {
	runes := []rune(str)
	
	curr := t;
	
	for r : runes {
		next := t.children[r]
		if next == nil {
			next = t.children[r] = &trie{word:false}
		}
		curr = next
	}
	curr.word = true
}

