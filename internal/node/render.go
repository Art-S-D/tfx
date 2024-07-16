package node

import (
	"strings"

	"github.com/Art-S-D/tfx/internal/style"
)

type Renderer interface {
	Render(node *Node) style.Str
}

type LineRenderer struct {
	Collapsed style.Str
	Expanded  style.Str
}

func (l *LineRenderer) Render(node *Node) style.Str {
	out := l.Collapsed
	if node.IsExpanded() {
		out = l.Expanded
	}
	return style.Spaces(int(node.Depth) * INDENT_WIDTH).UnSelectable().Concat(out)
}

type SensitiveRenderer struct {
	node *Node
}

func (s *SensitiveRenderer) Render(node *Node) style.Str {
	out := style.Sensitive("(sensitive)")
	if !node.Sensitive {
		out = s.node.Render(s.node)
	}
	return style.Spaces(int(node.Depth) * INDENT_WIDTH).UnSelectable().Concat(out)
}

type KeyValueRenderer struct {
	Key     string
	Value   Renderer
	Padding uint8
}

func (k *KeyValueRenderer) Render(node *Node) style.Str {
	out := k.Value.Render(node)
	out = style.TrimLeft(out)
	return style.Concat(
		style.Spaces(int(node.Depth)*INDENT_WIDTH).UnSelectable(),
		style.Key(k.Key).Selectable(),
		style.Spaces(int(k.Padding)),
		style.Default(" = ").UnSelectable(),
		out.UnSelectable(),
	)
}

// used for the root module
type NoRender struct{}

func (r *NoRender) Render(n *Node) style.Str {
	return nil
}

func (n *Node) View() style.Str {
	return n.Renderer.Render(n)
}

func RenderUnstyled(n *Node) string {
	var sb strings.Builder

	sb.WriteString(strings.Repeat(" ", int(n.Depth)*INDENT_WIDTH))

	view := n.View()
	sb.WriteString(style.NoColor.Render(view))

	for _, child := range n.Children() {
		sb.WriteRune('\n')
		sb.WriteString(RenderUnstyled(child))
	}
	return sb.String()

}
