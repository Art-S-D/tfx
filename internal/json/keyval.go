package json

import (
	"strings"

	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

type KeyVal struct {
	Key        string
	KeyPadding uint8
	Value      *render.Node
}

func (k *KeyVal) Children() []*render.Node {
	return k.Value.Children()
}

func (k *KeyVal) GenerateLines(node *render.Node) []render.Line {
	firstLine := render.Line{Indentation: node.Depth, PointsTo: node}
	firstLine.AddSelectable(style.Key(k.Key))
	firstLine.AddUnselectable(
		style.Default(strings.Repeat(" ", int(k.KeyPadding)-len(k.Key))),
		style.Default(" = "),
	)

	valueLines := k.Value.Lines()
	valueLines[0] = *firstLine.MergeWith(&valueLines[0])

	if len(valueLines) > 1 {
		valueLines[len(valueLines)-1].PointsTo = node
		valueLines[len(valueLines)-1].PointsToEnd = true
	}

	return valueLines
}
