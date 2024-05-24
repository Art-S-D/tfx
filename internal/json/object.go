package json

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
	"github.com/Art-S-D/tfx/internal/utils"
)

type jsonObject struct {
	render.BaseModel
	render.BaseCollapser
	value map[string]render.Model
}

func (o *jsonObject) Keys() []string {
	return utils.KeysOrdered(o.value)
}

func (b *jsonObject) Children() []render.Model {
	var out []render.Model
	keys := b.Keys()
	for _, k := range keys {
		out = append(out, b.value[k])
	}
	return out
}

func (o *jsonObject) View() []render.Line {
	firstLine := render.Line{PointsTo: o}

	if len(o.value) == 0 {
		firstLine.AddSelectable(style.Default("{}"))
		return []render.Line{firstLine}
	} else if !o.Expanded {
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

			lines := v.View()
			render.Indent(lines)
			out = append(out, lines...)
		}
		lastLine := render.Line{PointsTo: o, PointsToEnd: true}
		lastLine.AddSelectable(style.Default("}"))
		out = append(out, lastLine)
		return out
	}
}
