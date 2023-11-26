package trie

type node struct {
	charCode int32
	isWord   bool
	children [26]*node
}

func newNode(charCode int32) *node {
	return &node{
		charCode: charCode,
		isWord:   false,
		children: [26]*node{},
	}
}

func (n *node) childNode(charCode int32) *node {
	index := childNodeIndex(charCode)
	return n.children[index]
}

func (n *node) hasChildren() bool {
	for _, childNode := range n.children {
		if childNode != nil {
			return true
		}
	}
	return false
}

func childNodeIndex(charCode int32) int32 {
	return charCode - 'a'
}
