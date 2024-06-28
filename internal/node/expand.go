package node

func (n *Node) Expand() {
	n.isExpanded = true
}

func (n *Node) Collapse() {
	n.isExpanded = false
}
