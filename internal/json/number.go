package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
)

type jsonNumber struct {
	render.BaseModel
	value float64
}

func (n *jsonNumber) View(params render.ViewParams) []render.Token {
	line := render.Token{Theme: params.Theme, Indentation: params.Indentation, PointsTo: n, LineBreak: true}
	v := fmt.Sprintf("%.2f", n.value)
	line.AddSelectable(params.Theme.Number(v))
	return []render.Token{line}
}
