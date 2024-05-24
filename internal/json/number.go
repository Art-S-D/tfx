package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

type jsonNumber struct {
	render.BaseModel
	value float64
}

func (n *jsonNumber) View() []render.Line {
	line := render.Line{PointsTo: n}
	v := fmt.Sprintf("%.2f", n.value)
	line.AddSelectable(style.Number(v))
	return []render.Line{line}
}
