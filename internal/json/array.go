package json

import (
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

func (a *jsonArray) Lines(indent uint8) []*render.ScreenLine {

	firstLine := render.ScreenLine{Indentation: indent, PointsTo: a}

	if len(a.value) == 0 {
		firstLine.AddString(style.Default, "[]")
		return []*render.ScreenLine{&firstLine}
	}
	if !a.expanded {
		firstLine.AddString(style.Default, "[")
		firstLine.AddString(style.Preview, "...")
		firstLine.AddString(style.Default, "]")
		return []*render.ScreenLine{&firstLine}
	} else {
		var out []*render.ScreenLine
		firstLine.AddString(style.Default, "[")
		out = append(out, &firstLine)

		for i, v := range a.value {
			nextLines := v.Lines(indent + render.INDENT_WIDTH)

			if i < len(a.value)-1 {
				nextLines[len(nextLines)-1].AddUnSelectableString(style.Default, ",")
			}

			out = append(out, nextLines...)
		}

		lastLine := render.ScreenLine{Indentation: indent, PointsTo: a}
		lastLine.AddString(style.Default, "]")
		out = append(out, &lastLine)
		return out
	}
}
