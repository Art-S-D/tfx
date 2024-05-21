package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
)

type jsonString struct {
	render.BaseModel
	value string
}

func (s *jsonString) View(params render.ViewParams) []render.Token {
	line := render.Token{Theme: params.Theme, Indentation: params.Indentation, PointsTo: s, LineBreak: true}
	v := fmt.Sprintf("\"%s\"", s.value)
	line.AddSelectable(params.Theme.String(v))
	return []render.Token{line}
}
