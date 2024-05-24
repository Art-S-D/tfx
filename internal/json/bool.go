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

func (b *jsonBool) View() []render.Line {
	line := render.Line{PointsTo: b}
	line.AddSelectable(style.Boolean(fmt.Sprintf("%v", b.value)))
	return []render.Line{line}
}
