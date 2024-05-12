package render

import (
	"strings"

	"github.com/Art-S-D/tfx/internal/style"
	"github.com/charmbracelet/lipgloss"
)

type styledString struct {
	canBeSelected bool
	content       string
	style         lipgloss.Style
}

func (c styledString) render(cursor bool) string {
	if cursor && c.canBeSelected {
		return style.Cursor.Render(c.content)
	} else {
		return c.style.Render(c.content)
	}
}

type ScreenLine struct {
	content          []styledString
	Indentation      uint8
	PointsTo         Model // used to find the currently selected item
	PointsToModelEnd bool  // true if this the end of the PointsTo Model e.g: lines made of only } or ]
}

func LinesToString(lines []*ScreenLine) string {
	var sb strings.Builder
	for _, l := range lines {
		sb.WriteString(l.View(false))
	}
	return sb.String()
}

func (s *ScreenLine) AddString(style lipgloss.Style, str string) {
	s.content = append(
		s.content,
		styledString{
			canBeSelected: true,
			content:       str,
			style:         style,
		},
	)
}
func (s *ScreenLine) AddUnSelectableString(style lipgloss.Style, str string) {
	s.content = append(
		s.content,
		styledString{
			canBeSelected: false,
			content:       str,
			style:         style,
		},
	)
}

func (s *ScreenLine) indent() string {
	return strings.Repeat(" ", int(s.Indentation))
}

func (s *ScreenLine) View(onCursor bool) string {
	var sb strings.Builder
	sb.WriteString(s.indent())
	for _, s := range s.content {
		sb.WriteString(s.render(onCursor))
	}
	return sb.String()
}

func (s *ScreenLine) MergeWith(other *ScreenLine) {
	s.content = append(s.content, other.content...)
	s.PointsTo = other.PointsTo
}

func (s *ScreenLine) RemoveCursor() {
	for i := range s.content {
		s.content[i].canBeSelected = false
	}
}
