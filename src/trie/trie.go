package trie

import "fmt"

const (
	asciiOffset = 97
	alphabet    = 26
)

type trieNode struct {
	isWordEnd bool
	children  []*trieNode
}

type Trie struct {
	alphabet int
	root     *trieNode
}

func New() *Trie {
	return &Trie{alphabet, newNode(false, alphabet)}
}

func newNode(isWordEnd bool, size int) *trieNode {
	return &trieNode{isWordEnd, make([]*trieNode, size)}
}

func (t *Trie) Insert(s string) {
	if len(s) == 0 {
		return
	}

	p := t.root

	for i, r := range s {
		k := runeToNodeKey(r)
		node := p.children[k]
		isLast := i == len(s)-1

		if node == nil {
			n := newNode(isLast, t.alphabet)

			p.children[k] = n
			p = n
		}
	}
}

func (t *Trie) Search(s string) bool {
	if len(s) == 0 {
		return false
	}

	p := t.root

	for _, r := range s {
		k := runeToNodeKey(r)
		node := p.children[k]
		if node == nil {
			return false
		}

		p = node
	}

	return p.isWordEnd
}

func (t *Trie) StartsWith(s string) bool {
	if len(s) == 0 {
		return false
	}

	p := t.root

	for _, r := range s {
		k := runeToNodeKey(r)
		node := p.children[k]
		if node == nil {
			return false
		}

		p = node
	}

	return true
}

func (t *Trie) Print() {
	fmt.Println("Root:")
	print(t.root, 0)
}

func print(n *trieNode, level int) {
	var indent string

	for i := 0; i <= level; i++ {
		indent += "-"
	}

	for i, cn := range n.children {
		if cn != nil {
			fmt.Println(indent + keyToString(rune(i)))

			print(cn, level+1)
		}
	}
}

func runeToNodeKey(r rune) rune {
	return r - rune(asciiOffset)
}
func keyToString(r rune) string {
	return string(r + asciiOffset)
}
