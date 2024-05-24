package render

type Liner interface {
	GenerateLines(node *Node) []Line
}
type Childrener interface {
	Children() []*Node
}

type Node struct {
	Address  string
	Expanded bool
	Depth    uint8
	Liner    Liner
	Parent   *Node
	cache    []Line
}

func (n *Node) Children() []*Node {
	if childrener, ok := n.Liner.(Childrener); ok {
		return childrener.Children()
	}
	return []*Node{}
}

func (n *Node) ClearCache() {
	n.cache = nil
	if n.Parent != nil {
		n.Parent.ClearCache()
	}
}
func (n *Node) Lines() []Line {
	if n.cache == nil {
		n.cache = n.Liner.GenerateLines(n)
	}
	return n.cache
}

func (n *Node) Expand() {
	n.Expanded = true
	n.ClearCache()
}
func (n *Node) Collapse() {
	n.Expanded = false
	n.ClearCache()
}
