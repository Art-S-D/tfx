package json

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

type jsonNull struct {
	address string
}

func (n *jsonNull) Address() string {
	return n.address
}

func (n *jsonNull) Expand()   {}
func (n *jsonNull) Collapse() {}
func (n *jsonNull) ViewHeight() int {
	return 1
}

func (n *jsonNull) Children() []render.Model {
	return []render.Model{}
}

func (n *jsonNull) Lines(indent uint8) []*render.ScreenLine {
	line := render.ScreenLine{Indentation: indent, PointsTo: n}
	line.AddString(style.Null, "null")
	return []*render.ScreenLine{&line}
}
