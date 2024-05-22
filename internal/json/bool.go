package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

type jsonBool struct {
	render.BaseModel
	value bool
}

func (b *jsonBool) View(params render.ViewParams) []render.Line {
	line := render.Line{Indentation: params.Indentation, PointsTo: b}
	line.AddSelectable(style.Boolean(fmt.Sprintf("%v", b.value)))
	return []render.Line{line}
}
