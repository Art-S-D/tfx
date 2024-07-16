package node

func (n *Node) after(child *Node) *Node {
	for i, child2 := range n.Children() {
		if child == child2 {
			if i+1 < len(n.Children()) {
				return n.Children()[i+1]
			} else if n.Parent != nil {
				return n.Parent.after(n)
			} else {
				// should only happen for the very last node
				return nil
			}
		}
	}
	panic("after: child not found")
}

func (n *Node) Next() *Node {
	if !n.IsExpanded() {
		// if the node is collapsed, ignore the children
		return n.Parent.after(n)
	}
	if len(n.Children()) > 0 {
		return n.Children()[0]
	}
	return n.Parent.after(n)
}

// can be nil if there are no more siblings
func (n *Node) NextSibling() *Node {
	return n.Parent.after(n)
}

func (n *Node) LastChild() *Node {
	if len(n.Children()) == 0 || !n.IsExpanded() {
		return n
	} else {
		return n.Children()[len(n.Children())-1].LastChild()
	}
}

func (n *Node) before(child *Node) *Node {
	for i, child2 := range n.Children() {
		if child == child2 {
			if i > 0 {
				return n.Children()[i-1].LastChild()
			} else {
				// the node before the first child of n is n
				return n
			}
		}
	}
	panic("before: child not found")
}

func (n *Node) Previous() *Node {
	if n.Parent == nil {
		return nil
	} else {
		return n.Parent.before(n)
	}
}

func (n *Node) PreviousSibling() *Node {
	for i, child := range n.Parent.Children() {
		if child == n {
			if i == 0 {
				return n.Parent
			} else {
				return n.Parent.Children()[i-1] // no LastChild here
			}
		}
	}
	panic("PreviousSibling: child not found")
}
