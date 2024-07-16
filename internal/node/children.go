package node

func (n *Node) AppendChild(child *Node) {
	n.Childer.AppendChild(n, child)
}

type Childer interface {
	Children() []*Node
	AppendChild(node, child *Node)
	HasChild(child *Node) bool
}

type DefaultChilder struct {
	children []*Node
}

func (c *DefaultChilder) Children() []*Node {
	return c.children
}
func (c *DefaultChilder) AppendChild(node, child *Node) {
	c.children = append(c.children, child)
	child.Parent = node
}
func (c *DefaultChilder) HasChild(child *Node) bool {
	for _, c := range c.children {
		if c == child {
			return true
		}
	}
	return false
}

type NoChildren struct{}

func (n *NoChildren) Children() []*Node {
	return nil
}
func (n *NoChildren) AppendChild(_, _ *Node) {
	panic("trying to add child to a node without children")
}
func (n *NoChildren) HasChild(_ *Node) bool {
	return false
}
