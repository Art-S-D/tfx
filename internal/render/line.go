package render

import (
	"strings"

	"github.com/Art-S-D/tfx/internal/style"
	"github.com/charmbracelet/lipgloss"
)

// a line on the screen that can be rendered with the cursor or not
type Line struct {
	Theme       *style.Theme
	PointsTo    Model
	PointsToEnd bool // true if this the last line of an item, eg: }, ]

	// LineBreak    bool
	// Colon        bool
	// EndCursor    bool

	selectable   []lipgloss.Style
	unselectable []lipgloss.Style
}

// use a lipgloss.Style or a string
func (l *Line) AddUnselectable(s ...lipgloss.Style) {
	l.unselectable = append(l.unselectable, s...)
}

// use a lipgloss.Style or a string
func (l *Line) AddSelectable(s ...lipgloss.Style) {
	l.selectable = append(l.selectable, s...)
}

func (l *Line) String() string {
	var out strings.Builder
	for _, s := range l.selectable {
		out.WriteString(s.String())
	}
	for _, s := range l.unselectable {
		out.WriteString(s.String())
	}
	return out.String()
}

func (l *Line) renderLineElement(selected bool, style lipgloss.Style) string {
	if selected {
		return l.Theme.Cursor(style.Value()).String()
	} else {
		return style.String()
	}

}

func (l *Line) Render(selected bool) string {
	var out strings.Builder
	for _, s := range l.selectable {
		out.WriteString(l.renderLineElement(selected, s))
	}
	for _, s := range l.unselectable {
		out.WriteString(s.String())
	}

	// if l.Colon {
	// 	out.WriteRune(',')
	// }
	// if l.LineBreak {
	// 	out.WriteRune('\n')
	// }

	return out.String()
}
