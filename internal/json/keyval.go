package json

import (
	"fmt"
	"strings"

	"github.com/Art-S-D/tfx/internal/render"
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

func (k *KeyVal) View(params render.ViewParams) []render.Line {
	firstLine := render.Line{Theme: params.Theme, Indentation: params.Indentation, PointsTo: k}
	firstLine.AddSelectable(params.Theme.Key(fmt.Sprintf("\"%s\"", k.Key)))
	firstLine.AddUnselectable(
		params.Theme.Default(strings.Repeat(" ", int(k.KeyPadding)-len(k.Key))),
		params.Theme.Default(" = "),
	)

	valueLines := k.Value.View(params)
	valueLines[0] = *firstLine.MergeWith(&valueLines[0])

	if len(valueLines) > 1 {
		valueLines[len(valueLines)-1].PointsTo = k
		valueLines[len(valueLines)-1].PointsToEnd = true
	}

	return valueLines
}
