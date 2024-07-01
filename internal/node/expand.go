package node

func (n *Node) Expand() {
	n.isExpanded = true
}

func (n *Node) Collapse() {
	n.isExpanded = false
}

func (n *Node) ExpandRecursively() {
	n.Expand()
	for _, child := range n.children {
		child.ExpandRecursively()
	}
}
func (n *Node) CollapseRecursively() {
	n.Collapse()
	for _, child := range n.children {
		child.CollapseRecursively()
	}
}
