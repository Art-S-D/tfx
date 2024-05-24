package render

import (
	"strings"

	"github.com/Art-S-D/tfx/internal/style"
)

// a line on the screen that can be rendered with the cursor or not
type Line struct {
	PointsTo    *Node
	PointsToEnd bool // true if this the last line of an item, eg: }, ]

	Indentation uint8

	selectable   []style.Str
	unselectable []style.Str
}

func (l *Line) AddUnselectable(s ...style.Str) {
	l.unselectable = append(l.unselectable, s...)
}

func (l *Line) AddSelectable(s ...style.Str) {
	l.selectable = append(l.selectable, s...)
}

func (l *Line) MergeWith(other *Line) *Line {
	out := *l
	out.selectable = nil
	out.unselectable = nil
	out.selectable = append(out.selectable, l.selectable...)
	out.unselectable = append(out.unselectable, l.unselectable...)
	out.unselectable = append(out.unselectable, other.selectable...)
	out.unselectable = append(out.unselectable, other.unselectable...)
	return &out
}

func (l *Line) String() string {
	var out strings.Builder
	for range l.Indentation {
		out.WriteRune(' ')
	}

	for _, s := range l.selectable {
		out.WriteString(s.Value)
	}
	for _, s := range l.unselectable {
		out.WriteString(s.Value)
	}
	return out.String()
}

func (l *Line) Render(theme *style.Theme, selected bool) string {
	var out strings.Builder

	for range l.Indentation {
		out.WriteRune(' ')
	}

	for _, s := range l.selectable {
		out.WriteString(theme.RenderCursor(selected, s))
	}
	for _, s := range l.unselectable {
		out.WriteString(theme.Render(s))
	}

	return out.String()
}
