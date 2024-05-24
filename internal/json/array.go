package json

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

type jsonArray struct {
	value []*render.Node
}

func (a *jsonArray) Children() []*render.Node {
	return a.value
}

func (a *jsonArray) GenerateLines(node *render.Node) []render.Line {
	firstLine := render.Line{Indentation: node.Depth, PointsTo: node}

	if len(a.value) == 0 {
		firstLine.AddSelectable(style.Default("[]"))
		return []render.Line{firstLine}
	}
	if !node.Expanded {
		firstLine.AddSelectable(style.Default("["))
		firstLine.AddSelectable(style.Preview("..."))
		firstLine.AddSelectable(style.Default("]"))
		return []render.Line{firstLine}
	} else {
		firstLine.AddSelectable(style.Default("["))
		out := []render.Line{firstLine}

		for _, v := range a.value {
			lines := v.Lines()
			out = append(out, lines...)
		}

		lastLine := render.Line{Indentation: node.Depth, PointsTo: node, PointsToEnd: true}
		lastLine.AddSelectable(style.Default("]"))
		out = append(out, lastLine)
		return out
	}
}
