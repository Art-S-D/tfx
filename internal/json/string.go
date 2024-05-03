package json

import (
	"github.com/Art-S-D/tfview/internal/render"
	"github.com/Art-S-D/tfview/internal/style"
)

type jsonString struct {
	value   string
	address string
}

func (s *jsonString) Address() string {
	return s.address
}

func (s *jsonString) Expand()   {}
func (s *jsonString) Collapse() {}
func (s *jsonString) ViewHeight() int {
	return 1
}
func (s *jsonString) Selected(cursor int) (selected render.Model, cursorPosition int) {
	return s, 0
}
func (s *jsonString) View(params render.ViewParams) string {
	return style.RenderStyleOrCursor(params.Cursor, style.String, s.value)
}
