package json

import (
	"fmt"
	"strings"

	"github.com/Art-S-D/tfview/internal/render"
	"github.com/Art-S-D/tfview/internal/style"
	"github.com/Art-S-D/tfview/internal/utils"
)

type jsonObject struct {
	value    map[string]render.Model
	address  string
	expanded bool
}

func (o *jsonObject) Keys() []string {
	return utils.KeysOrdered(o.value)
}

func (o *jsonObject) Address() string {
	return o.address
}
func (o *jsonObject) Expand() {
	o.expanded = true
}
func (o *jsonObject) Collapse() {
	o.expanded = false
}
func (o *jsonObject) ViewHeight() int {
	if !o.expanded {
		return 1
	}
	// one for each curly brackets
	out := 2
	for _, v := range o.value {
		out += v.ViewHeight()
	}
	return out
}
func (o *jsonObject) Selected(cursor int) (selected render.Model, cursorPosition int) {
	if cursor == 0 {
		return o, 0
	}
	cursor -= 1
	for _, k := range o.Keys() {
		v := o.value[k]
		height := v.ViewHeight()
		if cursor < height {
			return v.Selected(cursor)
		} else {
			cursor -= height
		}
	}
	if cursor == 0 {
		return o, o.ViewHeight() - 1
	}
	panic(fmt.Sprintf("cursor out of bounds %d for %v of height %d", cursor, o, o.ViewHeight()))
}

func (o *jsonObject) View(params render.ViewParams) string {
	if len(o.value) == 0 {
		return style.RenderStyleOrCursor(params.Cursor, style.Default, "{}")
	}

	if !o.expanded {
		var sb strings.Builder
		sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Default, "{"))
		sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Preview, "..."))
		sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Default, "}"))
		return sb.String()
	} else {
		var sb strings.Builder
		sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Default, "{"))
		params.Cursor -= 1
		keys := o.Keys()
		for i, k := range keys {
			sb.WriteString("\n")
			sb.WriteString(style.Indented.Render(style.Key.Render(fmt.Sprintf("\"%v\"", k))))
			sb.WriteString("=")
			sb.WriteString(style.Indented.Render(o.value[k].View(params)))
			params.Cursor -= o.value[k].ViewHeight()
			if i < len(keys)-1 {
				sb.WriteRune(',')
			}
		}
		sb.WriteString("\n")
		sb.WriteString(style.RenderStyleOrCursor(params.Cursor, style.Default, "}"))
		return sb.String()
	}
}
