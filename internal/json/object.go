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

func (o *jsonObject) Lines(indent uint8) []*render.ScreenLine {
	firstLine := render.ScreenLine{Indentation: indent, PointsTo: o}

	if len(o.value) == 0 {
		firstLine.AddString(style.Default, "{}")
		return []*render.ScreenLine{&firstLine}
	} else if !o.expanded {
		firstLine.AddString(style.Default, "{")
		firstLine.AddString(style.Preview, "...")
		firstLine.AddString(style.Default, "}")
		return []*render.ScreenLine{&firstLine}
	} else {
		var out []*render.ScreenLine
		firstLine.AddString(style.Default, "{")
		out = append(out, &firstLine)

		keys := o.Keys()
		for i, k := range keys {
			v := o.value[k]

			line := render.ScreenLine{Indentation: indent + render.INDENT_WIDTH, PointsTo: v}
			quotedKey := fmt.Sprintf("\"%v\"", k)
			line.AddString(style.Key, quotedKey)
			line.AddUnSelectableString(style.Default, ": ")

			nextLines := v.Lines(indent + render.INDENT_WIDTH)
			nextLines[0].RemoveCursor()
			line.MergeWith(nextLines[0])

			if i < len(keys)-1 {
				nextLines[len(nextLines)-1].AddUnSelectableString(style.Default, ",")
			}

			out = append(out, &line)
			out = append(out, nextLines[1:]...)
		}

		lastLine := render.ScreenLine{Indentation: indent, PointsTo: o, PointsToModelEnd: true}
		lastLine.AddString(style.Default, "}")
		out = append(out, &lastLine)
		return out
	}
}
