package json

import (
	"fmt"

	"github.com/Art-S-D/tfview/internal/render"
	"github.com/Art-S-D/tfview/internal/style"
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
func (b *jsonBool) Selected(cursor int) (selected render.Model, cursorPosition int) {
	return b, 0
}

func (b *jsonBool) View(params render.ViewParams) string {
	v := fmt.Sprintf("%v", b.value)
	return style.RenderStyleOrCursor(params.Cursor, style.Boolean, v)
}
