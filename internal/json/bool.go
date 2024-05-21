package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
)

type jsonBool struct {
	render.BaseModel
	value bool
}

func (b *jsonBool) View(params render.ViewParams) []render.Token {
	line := render.Token{Theme: params.Theme, Indentation: params.Indentation, PointsTo: b, LineBreak: true}
	line.AddSelectable(params.Theme.Boolean(fmt.Sprintf("%v", b.value)))
	return []render.Token{line}
}
