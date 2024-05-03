package json

import (
	"fmt"
	"strings"

	"github.com/Art-S-D/tfview/internal/render"
	"github.com/Art-S-D/tfview/internal/style"
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
func (a *jsonArray) Selected(cursor int) (selected render.Model, cursorPosition int) {
	if cursor == 0 {
		return a, 0
	}
	for _, v := range a.value {
		height := v.ViewHeight()
		if cursor < height {
			return v, cursor
		} else {
			cursor -= height
		}
	}
	if cursor == 0 {
		return a, a.ViewHeight()
	}
	panic(fmt.Sprintf("cursor out of bounds %d for %v of height %d", cursor, a, a.ViewHeight()))
}

func (a *jsonArray) View(params render.ViewParams) string {
	if !a.expanded {
		var sb strings.Builder
		sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Default, "["))
		sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Preview, "..."))
		sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Default, "]"))
		return sb.String()
	} else {
		var sb strings.Builder
		sb.WriteRune('[')
		params.Cursor -= 1
		for i, v := range a.value {
			sb.WriteString("\n  ")
			sb.WriteString(v.View(params))
			params.Cursor -= v.ViewHeight()
			if i < len(a.value)-1 {
				sb.WriteRune(',')
			}
		}
		sb.WriteString("\n]")
		return sb.String()
	}
}
