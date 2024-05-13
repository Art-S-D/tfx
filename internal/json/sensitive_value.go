package json

import (
	"github.com/Art-S-D/tfx/internal/render"
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

func (v *SensitiveValue) ViewHeight() int {
	if !v.shown {
		return 1
	} else {
		return v.value.ViewHeight()
	}
}

func (v *SensitiveValue) Address() string {
	return v.value.Address()
}

// cannot be expanded or collapsed, only revealed
func (v *SensitiveValue) Expand() {
}
func (v *SensitiveValue) Collapse() {
}

func (v *SensitiveValue) Selected(cursor int) (selected render.Model, cursorPosition int) {
	if cursor == 0 && !v.shown {
		return v, 0
	} else {
		return v.value.Selected(cursor)
	}
}

func (v *SensitiveValue) Children() []render.Model {
	return v.value.Children()
}

func (v *SensitiveValue) View(params *render.ViewParams) string {
	builder := render.NewBuilder(params)
	if !v.shown {
		builder.AddString(params.Theme.Preview("(sensitive)"))
		return builder.String()
	} else {
		return v.value.View(params)
	}
}
