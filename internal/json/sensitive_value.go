package json

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

// assume it can only be revealed, not hidden afterwards
// once revealed it acts as if it does not exists
type SensitiveValue struct {
	value *render.Node
	shown bool
}

// this is a separate method from Expand so that 'expand all' doesn't reveal all sensitive values
// also it's probably better to have it on another key that expand for safety precautions
func (v *SensitiveValue) Reveal() {
	v.shown = true
	v.value.ClearCache()
}

func (v *SensitiveValue) Children() []*render.Node {
	return v.value.Children()
}

func (v *SensitiveValue) GenerateLines(node *render.Node) []render.Line {
	if !v.shown {
		line := render.Line{Indentation: node.Depth, PointsTo: node}
		line.AddSelectable(style.Preview("(sensitive)"))
		return []render.Line{line}
	} else {
		return v.value.Lines()
	}
}
