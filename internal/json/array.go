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

func (a *jsonArray) View(r *render.Renderer) {
	if len(a.value) == 0 {
		r.CursorWrite(style.Default, "[]")
		return
	}
	if !a.expanded {
		r.CursorWrite(style.Default, "[")
		r.CursorWrite(style.Preview, "...")
		r.CursorWrite(style.Default, "]")
		return
	} else {
		r.CursorWrite(style.Default, "[")
		r.IndentRight()

		for i, v := range a.value {
			r.NewLine()
			v.View(r)
			if i < len(a.value)-1 {
				r.Write(",")
			}
		}
		r.IndentLeft()
		r.NewLine()
		r.CursorWrite(style.Default, "]")
	}
}
