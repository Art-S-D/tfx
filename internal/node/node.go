package node

import "github.com/Art-S-D/tfx/internal/style"

const INDENT_WIDTH = 4

// a node represents one line in the screen
// it can have child nodes if it's an array, an object, a terraform module etc
type Node struct {
	Address string
	Depth   uint8
	Parent  *Node

	Sensitive bool

	Renderer
	Childer
	Expander
}

func (n *Node) Reveal() {
	n.Sensitive = false
}

func (n *Node) Indent() {
	n.IndentBy(1)
}

func (n *Node) IndentBy(by uint8) {
	n.Depth += by
	for _, child := range n.Children() {
		child.IndentBy(by)
	}
}

func Sensitive(n *Node) *Node {
	return &Node{
		Address:  n.Address,
		Depth:    n.Depth,
		Parent:   n.Parent,
		Renderer: &SensitiveRenderer{node: n},
		Childer:  &NoChildren{},
		Expander: &NonExpandable{},
	}
}

func StringNode(s string) *Node {
	str := style.Default(s).Selectable()
	return &Node{
		Renderer: &LineRenderer{
			Expanded:  str,
			Collapsed: str,
		},
		Childer:  &NoChildren{},
		Expander: &NonExpandable{},
	}
}
func StrNode(s style.Str) *Node {
	return &Node{
		Renderer: &LineRenderer{
			Expanded:  s,
			Collapsed: s,
		},
		Childer:  &NoChildren{},
		Expander: &NonExpandable{},
	}
}
