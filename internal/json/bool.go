package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

type jsonBool struct {
	value   bool
	address string
}

func (b *jsonBool) Address() string {
	return b.address
}
func (b *jsonBool) Expand()   {}
func (b *jsonBool) Collapse() {}
func (b *jsonBool) ViewHeight() int {
	return 1
}
func (b *jsonBool) Children() []render.Model {
	return []render.Model{}
}

func (b *jsonBool) Lines(indent uint8) []*render.ScreenLine {
	line := render.ScreenLine{Indentation: indent, PointsTo: b}
	line.AddString(style.Boolean, fmt.Sprintf("%v", b.value))
	return []*render.ScreenLine{&line}
}
