package node

func (n *Node) Expand() {
	if n.sensitive {
		return
	}
	n.Target().isExpanded = true
}

func (n *Node) Collapse() {
	n.Target().isExpanded = false
}

func (n *Node) ExpandRecursively() {
	if n.sensitive {
		return
	}
	n.Expand()
	for _, child := range n.children {
		child.ExpandRecursively()
	}
}

func (n *Node) CollapseRecursively() {
	if n.sensitive {
		return
	}
	n.Collapse()
	for _, child := range n.children {
		child.CollapseRecursively()
	}
}
