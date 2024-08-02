package node

import (
	"strings"

	"github.com/Art-S-D/tfx/internal/style"
)

// returns a model consisting of a single string
// useful for lines consisting soley of } or ]
func String(s string) *Node {
	line := style.Default(s).Selectable()
	out := &Node{
		collapsed: line,
		expanded:  line,
	}
	return out
}
func Str(s style.Str) *Node {
	out := &Node{
		collapsed: s,
		expanded:  s,
	}
	return out
}

func (n *Node) AddEndingColon() {
	n.collapsed = n.collapsed.Concat(style.Default(",").UnSelectable())

	if len(n.children) == 0 {
		n.expanded = n.expanded.Concat(style.Default(",").UnSelectable())
	} else {
		n.children[len(n.children)-1].AddEndingColon()
	}
}

// only one line
func (n *Node) View() style.Str {
	content := n.collapsed
	if n.isExpanded {
		content = n.expanded
	}
	if n.sensitive {
		content = style.Sensitive("(sensitive)")
	}
	if n.key != "" {
		content = style.Concat(
			style.Key(n.key).Selectable(),
			style.Spaces(int(n.keyPadding)),
			style.Default(" = ").UnSelectable(),
			content.UnSelectable(),
		)
	}
	return style.Spaces(int(n.depth) * INDENT_WIDTH).UnSelectable().Concat(content)
}

// raw render
func (n *Node) String() string {
	if n.sensitive {
		return "(sensitive)"
	}

	var sb strings.Builder

	sb.WriteString(strings.Repeat(" ", int(n.depth)*INDENT_WIDTH))

	if n.key != "" {
		sb.WriteString(n.key)
		sb.WriteString(strings.Repeat(" ", int(n.keyPadding)))
		sb.WriteString(" = ")
	}

	sb.WriteString(n.expanded.String())
	for _, child := range n.children {
		sb.WriteRune('\n')
		sb.WriteString(child.String())
	}
	return sb.String()
}
