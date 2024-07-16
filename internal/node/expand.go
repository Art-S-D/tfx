package node

type Expander interface {
	Expand()
	Collapse()
	IsExpanded() bool
	SetExpanded(expanded bool)
}

type DefaultExpander struct {
	expanded bool
}

func (e *DefaultExpander) Expand() {
	e.expanded = true
}
func (e *DefaultExpander) Collapse() {
	e.expanded = false
}
func (e *DefaultExpander) IsExpanded() bool {
	return e.expanded
}
func (e *DefaultExpander) SetExpanded(expanded bool) {
	e.expanded = expanded
}

type NonExpandable struct{}

func (e *NonExpandable) Expand()                   {}
func (e *NonExpandable) Collapse()                 {}
func (e *NonExpandable) SetExpanded(expanded bool) {}
func (e *NonExpandable) IsExpanded() bool {
	return true
}

func (n *Node) ExpandRecursively() {
	if n.Sensitive {
		return
	}
	n.Expand()
	for _, child := range n.Children() {
		child.ExpandRecursively()
	}
}

func (n *Node) CollapseRecursively() {
	if n.Sensitive {
		return
	}
	n.Collapse()
	for _, child := range n.Children() {
		child.CollapseRecursively()
	}
}
