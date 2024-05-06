package json

import (
	"fmt"

	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
	"github.com/Art-S-D/tfx/internal/utils"
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

func (o *jsonObject) View(r *render.Renderer) {
	if len(o.value) == 0 {
		r.CursorWrite(style.Default, "{}")
		return
	}

	if !o.expanded {
		r.CursorWrite(style.Default, "{")
		r.CursorWrite(style.Preview, "...")
		r.CursorWrite(style.Default, "}")
		return
	} else {
		r.CursorWrite(style.Default, "{")
		r.IndentRight()

		keys := o.Keys()
		for i, k := range keys {
			v := o.value[k]

			r.NewLine()
			quotedKey := fmt.Sprintf("\"%v\"", k)
			r.CursorWrite(style.Key, quotedKey)

			// this makes it so that only the key is selected instead of the key and the value
			// FIXME
			// vParams := params
			// if params.Cursor == 0 {
			// 	vParams.Cursor -= 1
			// }

			r.Write(": ")

			r.IndentRight()
			v.View(r)
			r.IndentLeft()

			if i < len(keys)-1 {
				r.Write(",")
			}
		}
		r.IndentLeft()
		r.NewLine()
		r.CursorWrite(style.Default, "}")
	}
}
