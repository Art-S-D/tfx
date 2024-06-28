package node

// When collapsing/expanding a node, if the node has children it will
// expand/collapse itself.
// If it has no children, the parent node will be used instead.
// The `target` function returns which node expand/collapse should target.
// func (n *Node) target() *Node {
// 	if n.parent == nil {
// 		return n
// 	}
// 	if len(n.children) == 0 {
// 		return n.parent
// 	}
// 	return n
// }

func (n *Node) Expand() {
	// target := n.target()
	// if target == n {
	// 	n.isExpanded = true
	// } else {
	// 	target.Expand()
	// }
	n.isExpanded = true
}

func (n *Node) Collapse() {
	// target := n.target()
	// if target == n {
	// 	n.isExpanded = false
	// } else {
	// 	target.Collapse()
	// }
	n.isExpanded = false
}
