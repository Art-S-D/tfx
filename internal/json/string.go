package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

type jsonString struct {
	value string
}

func (s *jsonString) GenerateLines(node *render.Node) []render.Line {
	line := render.Line{Indentation: node.Depth, PointsTo: node}
	v := fmt.Sprintf("\"%s\"", s.value)
	line.AddSelectable(style.String(v))
	return []render.Line{line}
}
