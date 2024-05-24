package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

type jsonBool struct {
	value bool
}

func (b *jsonBool) GenerateLines(node *render.Node) []render.Line {
	line := render.Line{Indentation: node.Depth, PointsTo: node}
	line.AddSelectable(style.Boolean(fmt.Sprintf("%v", b.value)))
	return []render.Line{line}
}
