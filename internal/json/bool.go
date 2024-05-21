package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
)

type jsonBool struct {
	render.BaseModel
	value bool
}

func (b *jsonBool) View(params render.ViewParams) []render.Line {
	line := render.Line{Theme: params.Theme, PointsTo: b}
	line.AddSelectable(params.Theme.Boolean(fmt.Sprintf("%v", b.value)))
	return []render.Line{line}
}
