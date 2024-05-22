package json

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

type jsonArray struct {
	render.BaseModel
	render.BaseCollapser
	value []render.Model
}

func (a *jsonArray) Children() []render.Model {
	return a.value
}

func (a *jsonArray) View(params render.ViewParams) []render.Line {
	firstLine := render.Line{Indentation: params.Indentation, PointsTo: a}

	if len(a.value) == 0 {
		firstLine.AddSelectable(style.Default("[]"))
		return []render.Line{firstLine}
	}
	if !a.Expanded {
		firstLine.AddSelectable(style.Default("["))
		firstLine.AddSelectable(style.Preview("..."))
		firstLine.AddSelectable(style.Default("]"))
		return []render.Line{firstLine}
	} else {
		firstLine.AddSelectable(style.Default("["))
		out := []render.Line{firstLine}

		for _, v := range a.value {
			lines := v.View(params.IndentedRight())
			out = append(out, lines...)
		}

		lastLine := render.Line{Indentation: params.Indentation, PointsTo: a, PointsToEnd: true}
		lastLine.AddSelectable(style.Default("]"))
		out = append(out, lastLine)
		return out
	}
}
