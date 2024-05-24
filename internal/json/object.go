package json

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
	"github.com/Art-S-D/tfx/internal/utils"
)

type jsonObject struct {
	value map[string]*render.Node
}

func (o *jsonObject) Keys() []string {
	return utils.KeysOrdered(o.value)
}

func (b *jsonObject) Children() []*render.Node {
	var out []*render.Node
	keys := b.Keys()
	for _, k := range keys {
		out = append(out, b.value[k])
	}
	return out
}

func (o *jsonObject) GenerateLines(node *render.Node) []render.Line {
	firstLine := render.Line{Indentation: node.Depth, PointsTo: node}

	if len(o.value) == 0 {
		firstLine.AddSelectable(style.Default("{}"))
		return []render.Line{firstLine}
	} else if !node.Expanded {
		firstLine.AddSelectable(style.Default("{"))
		firstLine.AddSelectable(style.Preview("..."))
		firstLine.AddSelectable(style.Default("}"))
		return []render.Line{firstLine}
	} else {
		firstLine.AddSelectable(style.Default("{"))
		out := []render.Line{firstLine}

		keys := o.Keys()
		for _, k := range keys {
			v := o.value[k]
			lines := v.Lines()
			out = append(out, lines...)
		}
		lastLine := render.Line{Indentation: node.Depth, PointsTo: node, PointsToEnd: true}
		lastLine.AddSelectable(style.Default("}"))
		out = append(out, lastLine)
		return out
	}
}
