package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

type jsonNumber struct {
	value float64
}

func (n *jsonNumber) GenerateLines(node *render.Node) []render.Line {
	line := render.Line{Indentation: node.Depth, PointsTo: node}
	v := fmt.Sprintf("%.2f", n.value)
	line.AddSelectable(style.Number(v))
	return []render.Line{line}
}
