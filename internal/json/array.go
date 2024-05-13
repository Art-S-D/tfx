package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
)

type jsonArray struct {
	value    []render.Model
	address  string
	expanded bool
}

func (a *jsonArray) Address() string {
	return a.address
}

func (a *jsonArray) Expand() {
	a.expanded = true
}
func (a *jsonArray) Collapse() {
	a.expanded = false
}
func (a *jsonArray) ViewHeight() int {
	if !a.expanded {
		return 1
	}
	// one for each square brackets
	out := 2
	for _, v := range a.value {
		out += v.ViewHeight()
	}
	return out
}
func (a *jsonArray) Children() []render.Model {
	return a.value
}
func (a *jsonArray) Selected(cursor int) (selected render.Model, cursorPosition int) {
	if cursor == 0 {
		return a, 0
	}
	cursor -= 1
	for _, v := range a.value {
		height := v.ViewHeight()
		if cursor < height {
			return v.Selected(cursor)
		} else {
			cursor -= height
		}
	}
	if cursor == 0 {
		return a, a.ViewHeight() - 1
	}
	panic(fmt.Sprintf("cursor out of bounds %d for %v of height %d", cursor, a, a.ViewHeight()))
}

func (a *jsonArray) View(params *render.ViewParams) string {
	builder := render.NewBuilder(params)

	if len(a.value) == 0 {
		builder.AddString(params.Theme.Default("[]"))
		return builder.String()
	}
	if !a.expanded {
		builder.AddString(params.Theme.Default("["))
		builder.AddString(params.Theme.Preview("..."))
		builder.AddString(params.Theme.Default("]"))
		return builder.String()
	} else {
		builder.AddString(params.Theme.Default("["))
		params.IndentRight()

		for i, v := range a.value {
			builder.InsertNewLine()
			params.NextLine()

			builder.WriteString(v.View(params))
			if i < len(a.value)-1 {
				builder.WriteString(",")
			}
		}

		params.IndentLeft()
		builder.InsertNewLine()
		params.NextLine()
		builder.AddString(params.Theme.Default("]"))
		return builder.String()
	}
}
