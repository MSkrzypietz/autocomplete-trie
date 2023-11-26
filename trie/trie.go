package trie

type Trie struct {
	head *node
}

func NewTrie() *Trie {
	return &Trie{head: newNode(-1)}
}

func (t *Trie) Insert(item string) {
	node := t.head
	for _, ch := range item {
		index := childNodeIndex(ch)
		if node.children[index] != nil {
			node = node.children[index]
		} else {
			node.children[index] = newNode(ch)
			node = node.children[index]
		}
	}
	node.isWord = true
}

func (t *Trie) Delete(item string) {
	if len(item) > 0 {
		innerDelete(t.head, item)
	}
}

func innerDelete(node *node, item string) bool {
	if node == nil {
		return false
	}
	if len(item) == 0 {
		node.isWord = false
		return true
	}

	childNode := node.childNode(int32(item[0]))
	found := innerDelete(childNode, item[1:])
	if !found || childNode.hasChildren() || childNode.isWord {
		return false
	}

	childNode = nil
	return true
}

func (t *Trie) Find(partial string) []string {
	node := t.head
	for _, ch := range partial {
		childNode := node.childNode(ch)
		if childNode != nil {
			node = childNode
		} else {
			return []string{}
		}
	}
	return dfs(node, partial[:len(partial)-1])
}

func dfs(node *node, partial string) []string {
	var result []string
	if node == nil {
		return result
	}

	partial = partial + string(node.charCode)
	if node.isWord {
		result = append(result, partial)
	}

	for _, childNode := range node.children {
		result = append(result, dfs(childNode, partial)...)
	}

	return result
}
