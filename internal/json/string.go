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

func (s *jsonString) Children() []render.Model {
	return []render.Model{}
}

func (s *jsonString) Lines(indent uint8) []*render.ScreenLine {
	line := render.ScreenLine{Indentation: indent, PointsTo: s}
	line.AddString(style.String, fmt.Sprintf("\"%s\"", s.value))
	return []*render.ScreenLine{&line}
}
