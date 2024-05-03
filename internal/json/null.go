package json

import (
	"github.com/Art-S-D/tfview/internal/render"
	"github.com/Art-S-D/tfview/internal/style"
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
func (s *jsonNull) View(params render.ViewParams) string {
	return style.RenderStyleOrCursor(params.Cursor, style.Null, "null")
}
