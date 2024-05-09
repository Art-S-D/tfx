package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
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
func (s *jsonString) Children() []render.Model {
	return []render.Model{}
}
func (s *jsonString) View(params *render.ViewParams) string {
	builder := render.NewBuilder(params)
	v := fmt.Sprintf("\"%s\"", s.value)
	builder.WriteStyleOrCursor(style.String, v)
	return builder.String()
}
