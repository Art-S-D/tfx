package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
)

type jsonNumber struct {
	render.BaseModel
	value float64
}

func (n *jsonNumber) View(params render.ViewParams) []render.Line {
	line := render.Line{Theme: params.Theme, PointsTo: n}
	v := fmt.Sprintf("%.2f", n.value)
	line.AddSelectable(params.Theme.Number(v))
	return []render.Line{line}
}
