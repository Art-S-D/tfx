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
func (n *jsonNull) Selected(cursor int) (selected render.Model, cursorPosition int) {
	return n, 0
}
func (n *jsonNull) Children() []render.Model {
	return []render.Model{}
}
func (s *jsonNull) View(r *render.Renderer) {
	r.CursorWrite(style.Null, "null")
}
