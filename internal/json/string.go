package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
)

type jsonString struct {
	render.BaseModel
	value string
}

func (s *jsonString) View(params render.ViewParams) []render.Line {
	line := render.Line{Theme: params.Theme, PointsTo: s}
	v := fmt.Sprintf("\"%s\"", s.value)
	line.AddSelectable(params.Theme.String(v))
	return []render.Line{line}
}
