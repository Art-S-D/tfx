package json

import (
	"github.com/Art-S-D/tfx/internal/render"
)

type jsonNull struct {
	render.BaseModel
}

func (n *jsonNull) View(params render.ViewParams) []render.Token {
	line := render.Token{Theme: params.Theme, Indentation: params.Indentation, PointsTo: n, LineBreak: true}
	line.AddSelectable(params.Theme.Null("null"))
	return []render.Token{line}
}
