package educative

type trieNode struct {
	childrens [26]*trieNode
	wordEnds  bool
	word      string
}

type Trie struct {
	root *trieNode
}

func Constructor() Trie {
	return Trie{
		root: new(trieNode),
	}
}

func (t *Trie) Insert(word string) {
	current := t.root

	for _, wr := range word {
		index := wr - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = new(trieNode)
		}
		current = current.childrens[index]
	}

	current.wordEnds = true
	current.word = word
}

func (t *Trie) Search(word string) bool {
	current := t.root
	for _, wr := range word {
		index := wr - 'a'
		if current.childrens[index] == nil {
			return false
		}
		current = current.childrens[index]
	}

	return current.wordEnds
}

func (t *Trie) StartsWith(prefix string) bool {
	current := t.root
	for _, wr := range prefix {
		index := wr - 'a'
		if current.childrens[index] == nil {
			return false
		}
		current = current.childrens[index]
	}
	return true
}
