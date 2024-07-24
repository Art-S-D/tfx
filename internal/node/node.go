package node

import (
	"slices"

	"github.com/Art-S-D/tfx/internal/style"
)

const INDENT_WIDTH = 4

// a node represents one line in the screen
// it can have child nodes if it's an array, an object, a terraform module etc
type Node struct {
	address    string
	isExpanded bool
	depth      uint8

	// only represents the first line of the Node
	// the reste of the node if made of its children
	// including a potential last line such as ] or }
	expanded  style.Str
	collapsed style.Str

	// the value is sensitive and will be shown as `(sensitive)`
	sensitive bool

	// if the parent is an object, this Node is a key/value
	// this should be reflected in the expanded / collapsed line
	key        string
	keyPadding uint8

	parent   *Node
	children []*Node
}

func (n *Node) Address() string             { return n.address }
func (n *Node) Parent() *Node               { return n.parent }
func (n *Node) IsExpanded() bool            { return n.isExpanded }
func (n *Node) Depth() uint8                { return n.depth }
func (n *Node) SetParent(m *Node)           { n.parent = m }
func (n *Node) SetDepth(depth uint8)        { n.depth = depth }
func (n *Node) SetExpanded(line style.Str)  { n.expanded = line }
func (n *Node) SetCollapsed(line style.Str) { n.collapsed = line }
func (n *Node) SetAddress(addr string)      { n.address = addr }
func (n *Node) SetSensitive(sensitive bool) { n.sensitive = sensitive }
func (n *Node) Children() []*Node           { return n.children }

func (n *Node) IsRootModule() bool {
	return n.address == ""
}

func (n *Node) HasChild(child *Node) bool {
	return slices.Contains(n.children, child)
}

func (n *Node) Reveal() {
	n.sensitive = false
}

func (n *Node) SetKey(key string, padding uint8) {
	n.key = key
	n.keyPadding = padding
}
func (n *Node) AppendChild(child *Node) {
	child.parent = n
	n.children = append(n.children, child)
}

func (n *Node) IncreaseDepth() {
	n.depth += 1
	for _, child := range n.children {
		child.IncreaseDepth()
	}
}
func (n *Node) IncreaseDepthBy(by uint8) {
	n.depth += by
	for _, child := range n.children {
		child.IncreaseDepthBy(by)
	}
}

func (n *Node) Clone() *Node {
	out := *n
	out.parent = nil
	out.children = []*Node{}
	for _, child := range n.children {
		out.AppendChild(child.Clone())
	}
	return &out
}
