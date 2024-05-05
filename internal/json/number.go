package json

import (
	"fmt"

	"github.com/Art-S-D/tfview/internal/render"
	"github.com/Art-S-D/tfview/internal/style"
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
func (n *jsonNumber) Selected(cursor int) (selected render.Model, cursorPosition int) {
	return n, 0
}

func (n *jsonNumber) View(r *render.Renderer) {
	v := fmt.Sprintf("%.2f", n.value)
	r.CursorWrite(style.Number, v)
}
