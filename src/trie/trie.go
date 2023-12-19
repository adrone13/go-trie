package trie

import "fmt"

const (
	asciiOffset = 97
	alphabet    = 26
)

type TrieNode struct {
	IsWordEnd bool
	Children  []*TrieNode
}

type Trie struct {
	alphabet int
	root     *TrieNode
}

func New() *Trie {
	return &Trie{alphabet, newNode(false, alphabet)}
}

func newNode(isWordEnd bool, size int) *TrieNode {
	return &TrieNode{isWordEnd, make([]*TrieNode, size)}
}

func (t *Trie) Insert(s string) {
	if len(s) == 0 {
		return
	}

	p := t.root

	for i, r := range s {
		k := runeToNodeKey(r)
		node := p.Children[k]
		isLast := i == len(s)-1

		if node == nil {
			n := newNode(isLast, t.alphabet)

			p.Children[k] = n
			p = n
		}
	}
}

func (t *Trie) Search(s string) bool {
	if len(s) == 0 {
		return false
	}

	p := t.root

	for i, r := range s {
		k := runeToNodeKey(r)
		node := p.Children[k]
		if node == nil {
			return false
		}

		isLast := i == len(s)-1
		if isLast && !node.IsWordEnd {
			return false
		}

		p = node
	}

	return true
}

func (t *Trie) StartsWith(s string) bool {
	if len(s) == 0 {
		return false
	}

	p := t.root

	for _, r := range s {
		k := runeToNodeKey(r)
		node := p.Children[k]
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

func print(n *TrieNode, level int) {
	var indent string

	for i := 0; i <= level; i++ {
		indent += "-"
	}

	for i, cn := range n.Children {
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
