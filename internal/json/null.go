package json

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

type jsonNull struct {
}

func (n *jsonNull) GenerateLines(node *render.Node) []render.Line {
	line := render.Line{Indentation: node.Depth, PointsTo: node}
	line.AddSelectable(style.Null("null"))
	return []render.Line{line}
}
