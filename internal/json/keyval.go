package json

import (
	"strings"

	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

type KeyVal struct {
	Key        string
	KeyPadding uint8
	Value      render.Model
}

func (k *KeyVal) Address() string {
	return k.Value.Address()
}

func (k *KeyVal) Expand() {
	if expand, ok := k.Value.(render.Collapser); ok {
		expand.Expand()
	}
}
func (k *KeyVal) Collapse() {
	if expand, ok := k.Value.(render.Collapser); ok {
		expand.Collapse()
	}
}

func (k *KeyVal) Children() []render.Model {
	if child, ok := k.Value.(render.Childrener); ok {
		return child.Children()
	}
	return []render.Model{}
}

func (k *KeyVal) View() []render.Line {
	firstLine := render.Line{PointsTo: k}
	firstLine.AddSelectable(style.Key(k.Key))
	firstLine.AddUnselectable(
		style.Default(strings.Repeat(" ", int(k.KeyPadding)-len(k.Key))),
		style.Default(" = "),
	)

	valueLines := k.Value.View()
	valueLines[0] = *firstLine.MergeWith(&valueLines[0])

	if len(valueLines) > 1 {
		valueLines[len(valueLines)-1].PointsTo = k
		valueLines[len(valueLines)-1].PointsToEnd = true
	}

	return valueLines
}
