package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
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
		builder.WriteStyleOrCursor(style.Default, "[]")
		return builder.String()
	}
	if !a.expanded {
		builder.WriteStyleOrCursor(style.Default, "[")
		builder.WriteStyleOrCursor(style.Preview, "...")
		builder.WriteStyleOrCursor(style.Default, "]")
		return builder.String()
	} else {
		builder.WriteStyleOrCursor(style.Default, "[")

		for i, v := range a.value {
			params.NextLine()
			builder.InsertNewLine()
			builder.WriteString(v.View(params.IndentedRight()))
			if i < len(a.value)-1 {
				builder.WriteString(",")
			}
		}

		params.NextLine()
		builder.InsertNewLine()
		builder.WriteStyleOrCursor(style.Default, "]")
		return builder.String()
	}
}
