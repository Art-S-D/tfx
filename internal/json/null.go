package json

import (
	"github.com/Art-S-D/tfx/internal/render"
)

type jsonNull struct {
	render.BaseModel
}

func (n *jsonNull) View(params render.ViewParams) []render.Line {
	line := render.Line{Theme: params.Theme, PointsTo: n}
	line.AddSelectable(params.Theme.Null("null"))
	return []render.Line{line}
}
