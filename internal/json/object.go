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

func (o *jsonObject) View(params render.ViewParams) []render.Token {
	firstLine := render.Token{Theme: params.Theme, Indentation: params.Indentation, PointsTo: o, LineBreak: true}

	if len(o.value) == 0 {
		firstLine.AddSelectable(params.Theme.Default("{}"))
		return []render.Token{firstLine}
	} else if !o.Expanded {
		firstLine.AddSelectable(params.Theme.Default("{"))
		firstLine.AddSelectable(params.Theme.Preview("..."))
		firstLine.AddSelectable(params.Theme.Default("}"))
		return []render.Token{firstLine}
	} else {
		firstLine.AddSelectable(params.Theme.Default("{"))
		out := []render.Token{firstLine}

		keys := o.Keys()
		for _, k := range keys {
			v := o.value[k]

			line := render.Token{Theme: params.Theme, Indentation: params.IndentedRight().Indentation, PointsTo: v, EndCursor: true}

			quotedKey := fmt.Sprintf("\"%v\"", k)
			line.AddSelectable(params.Theme.Key(quotedKey))
			line.AddUnselectable(params.Theme.Default(": "))
			out = append(out, line)

			lines := v.View(params.IndentedRight())
			out = append(out, lines...)
		}
		lastLine := render.Token{Theme: params.Theme, Indentation: params.Indentation, PointsTo: o, PointsToEnd: true, LineBreak: true}
		lastLine.AddSelectable(params.Theme.Default("}"))
		out = append(out, lastLine)
		return out
	}
}
