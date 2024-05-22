package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

type jsonString struct {
	render.BaseModel
	value string
}

func (s *jsonString) View(params render.ViewParams) []render.Line {
	line := render.Line{Indentation: params.Indentation, PointsTo: s}
	v := fmt.Sprintf("\"%s\"", s.value)
	line.AddSelectable(style.String(v))
	return []render.Line{line}
}
