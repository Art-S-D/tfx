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
func (b *jsonObject) Children() []render.Model {
	var out []render.Model
	keys := b.Keys()
	for _, k := range keys {
		out = append(out, b.value[k])
	}
	return out
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

func (o *jsonObject) View(params *render.ViewParams) string {
	builder := render.NewBuilder(params)

	if len(o.value) == 0 {
		builder.WriteStyleOrCursor(style.Default, "{}")
		return builder.String()
	}

	if !o.expanded {
		builder.WriteStyleOrCursor(style.Default, "{")
		builder.WriteStyleOrCursor(style.Preview, "...")
		builder.WriteStyleOrCursor(style.Default, "}")
		return builder.String()
	} else {
		builder.WriteStyleOrCursor(style.Default, "{")
		params.IndentRight()

		keys := o.Keys()
		for i, k := range keys {
			v := o.value[k]

			builder.InsertNewLine()
			params.NextLine()
			quotedKey := fmt.Sprintf("\"%v\"", k)
			builder.WriteStyleOrCursor(style.Key, quotedKey)

			params.EndCursorForCurrentLine()
			builder.WriteString(": ")
			builder.WriteString(v.View(params))

			if i < len(keys)-1 {
				builder.WriteString(",")
			}
		}
		params.IndentLeft()
		builder.InsertNewLine()
		params.NextLine()
		builder.WriteStyleOrCursor(style.Default, "}")
		return builder.String()
	}
}
