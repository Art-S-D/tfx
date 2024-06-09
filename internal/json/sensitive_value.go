package json

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

// assume it can only be revealed, not hidden afterwards
// once revealed it acts as if it does not exists
type SensitiveValue struct {
	value render.Model
	shown bool
}

// this is a separate method from Expand so that 'expand all' doesn't reveal all sensitive values
// also it's probably better to have it on another key that expand for safety precautions
func (v *SensitiveValue) Reveal() {
	v.shown = true
}
func (s *SensitiveValue) Address() string {
	return s.value.Address()
}

func (v *SensitiveValue) Children() []render.Model {
	return v.value.Children()
}

func (v *SensitiveValue) Expand() {
	if !v.shown {
		return
	}
	v.value.Expand()
}
func (v *SensitiveValue) Collapse() {
	if !v.shown {
		return
	}
	v.value.Collapse()
}
func (v *SensitiveValue) IsCollapsed() bool {
	if !v.shown {
		return true
	}
	return v.value.IsCollapsed()
}

func (v *SensitiveValue) View() []render.Line {
	if !v.shown {
		line := render.Line{PointsTo: v}
		line.AddSelectable(style.Preview("(sensitive)"))
		return []render.Line{line}
	} else {
		return v.value.View()
	}
}
