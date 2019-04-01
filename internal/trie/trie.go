package trie

type Node struct {
	children map[rune]*Node
	isTerminal bool
	value rune
}

func New() *Node {
	return &Node{
		children: make(map[rune]*Node),
	}
}

func (node *Node) Insert(word string) {
	cur := node
	for _, char := range []rune(word) {
		letter := char
		child, ok := cur.children[letter]
		if ok {
			cur = child
		} else {
			child := New()
			child.value = letter
			cur.children[letter] = child
			cur = child
		}
	}
	cur.isTerminal = true
}

func (node *Node) Exists(word string) (prefixExists, wordExists bool) {
	if word == "" {
		return false, false
	}
	cur := node
	for _, char := range []rune(word) {
		letter := char

		child, ok := cur.children[letter]
		if ok {
			cur = child
		} else {
			return false, false
		}
	}
	return true, cur.isTerminal
}
