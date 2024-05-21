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

func (v *SensitiveValue) Address() string {
	return v.value.Address()
}

func (v *SensitiveValue) Children() []render.Model {
	if childrener, ok := v.value.(render.Childrener); ok {
		return childrener.Children()
	} else {
		return []render.Model{}
	}
}

func (v *SensitiveValue) View(params render.ViewParams) []render.Token {
	if !v.shown {
		line := render.Token{Theme: params.Theme, Indentation: params.Indentation, PointsTo: v, LineBreak: true}
		line.AddSelectable(params.Theme.Preview("(sensitive)"))
		return []render.Token{line}
	} else {
		return v.value.View(params)
	}
}
