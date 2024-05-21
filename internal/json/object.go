package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
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

func (o *jsonObject) View(params render.ViewParams) []render.Line {
	firstLine := render.Line{Theme: params.Theme, PointsTo: o}

	if len(o.value) == 0 {
		firstLine.AddSelectable(params.Theme.Default("{}"))
		return []render.Line{firstLine}
	} else if !o.Expanded {
		firstLine.AddSelectable(params.Theme.Default("{"))
		firstLine.AddSelectable(params.Theme.Preview("..."))
		firstLine.AddSelectable(params.Theme.Default("}"))
		return []render.Line{firstLine}
	} else {
		firstLine.AddSelectable(params.Theme.Default("{"))
		out := []render.Line{firstLine}

		keys := o.Keys()
		for _, k := range keys {
			v := o.value[k]

			line := render.Line{Theme: params.Theme, PointsTo: v}

			quotedKey := fmt.Sprintf("\"%v\"", k)
			line.AddSelectable(params.Theme.Key(quotedKey))
			line.AddUnselectable(params.Theme.Default(": "))
			out = append(out, line)

			lines := v.View(params.IndentedRight())
			out = append(out, lines...)
		}
		lastLine := render.Line{Theme: params.Theme, PointsTo: o, PointsToEnd: true}
		lastLine.AddSelectable(params.Theme.Default("}"))
		out = append(out, lastLine)
		return out
	}
}
