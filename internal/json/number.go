package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

type jsonNumber struct {
	value   float64
	address string
}

func (n *jsonNumber) Address() string {
	return n.address
}

func (n *jsonNumber) Expand()   {}
func (n *jsonNumber) Collapse() {}
func (n *jsonNumber) ViewHeight() int {
	return 1
}

func (n *jsonNumber) Children() []render.Model {
	return []render.Model{}
}

func (n *jsonNumber) Lines(indent uint8) []*render.ScreenLine {
	line := render.ScreenLine{Indentation: indent, PointsTo: n}
	line.AddString(style.Number, fmt.Sprintf("%.2f", n.value))
	return []*render.ScreenLine{&line}
}
