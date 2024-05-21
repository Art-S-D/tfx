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

func (a *jsonArray) View(params render.ViewParams) []render.Token {
	firstLine := render.Token{Theme: params.Theme, Indentation: params.Indentation, PointsTo: a, LineBreak: true}

	if len(a.value) == 0 {
		firstLine.AddSelectable(params.Theme.Default("[]"))
		return []render.Token{firstLine}
	}
	if !a.Expanded {
		firstLine.AddSelectable(
			params.Theme.Default("["),
			params.Theme.Preview("..."),
			params.Theme.Default("]"),
		)
		return []render.Token{firstLine}
	} else {
		firstLine.AddSelectable(params.Theme.Default("["))
		out := []render.Token{firstLine}

		for _, v := range a.value {
			lines := v.View(params.IndentedRight())
			out = append(out, lines...)
		}

		lastLine := render.Token{Theme: params.Theme, Indentation: params.Indentation, PointsTo: a, PointsToEnd: true, LineBreak: true}
		lastLine.AddSelectable(params.Theme.Default("]"))
		out = append(out, lastLine)
		return out
	}
}
