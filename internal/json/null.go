package json

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

type jsonNull struct {
	render.BaseModel
}

func (n *jsonNull) View(params render.ViewParams) []render.Line {
	line := render.Line{Indentation: params.Indentation, PointsTo: n}
	line.AddSelectable(style.Null("null"))
	return []render.Line{line}
}
