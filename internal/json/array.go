package json

import (
	"github.com/Art-S-D/tfx/internal/render"
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
	firstLine := render.Line{Theme: params.Theme, Indentation: params.Indentation, PointsTo: a}

	if len(a.value) == 0 {
		firstLine.AddSelectable(params.Theme.Default("[]"))
		return []render.Line{firstLine}
	}
	if !a.Expanded {
		firstLine.AddSelectable(params.Theme.Default("["))
		firstLine.AddSelectable(params.Theme.Preview("..."))
		firstLine.AddSelectable(params.Theme.Default("]"))
		return []render.Line{firstLine}
	} else {
		firstLine.AddSelectable(params.Theme.Default("["))
		out := []render.Line{firstLine}

		for _, v := range a.value {
			lines := v.View(params.IndentedRight())
			out = append(out, lines...)
		}

		lastLine := render.Line{Theme: params.Theme, Indentation: params.Indentation, PointsTo: a, PointsToEnd: true}
		lastLine.AddSelectable(params.Theme.Default("]"))
		out = append(out, lastLine)
		return out
	}
}
