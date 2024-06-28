package node

func (n *Node) after(child *Node) *Node {
	for i, child2 := range n.children {
		if child == child2 {
			if i+1 < len(n.children) {
				return n.children[i+1]
			} else if n.parent != nil {
				return n.parent.after(n)
			} else {
				// should only happen for the very last node
				return nil
			}
		}
	}
	panic("child not found")
}

func (n *Node) Next() *Node {
	if !n.isExpanded {
		// if the node is collapsed, ignore the children
		return n.parent.after(n)
	}
	if len(n.children) > 0 {
		return n.children[0]
	}
	return n.parent.after(n)
}

func (n *Node) LastChild() *Node {
	if len(n.children) == 0 || !n.isExpanded {
		return n
	} else {
		return n.children[len(n.children)-1].LastChild()
	}
}

func (n *Node) before(child *Node) *Node {
	for i, child2 := range n.children {
		if child == child2 {
			if i > 0 {
				return n.children[i-1].LastChild()
			} else {
				// the node before the first child of n is n
				return n
			}
		}
	}
	panic("child not found")
}

func (n *Node) Previous() *Node {
	if n.parent == nil {
		return nil
	} else {
		return n.parent.before(n)
	}
}
