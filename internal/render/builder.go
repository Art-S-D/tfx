package render

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// a wrapper around strings.Builder that can manage rendering indentation and the cursor
type Builder struct {
	builder strings.Builder
	params  *ViewParams
}

func NewBuilder(params *ViewParams) *Builder {
	var builder strings.Builder
	return &Builder{
		params:  params,
		builder: builder,
	}
}

func (r *Builder) AddUnSelectableString(s lipgloss.Style) {
	if r.params.CurrentLineIsInView() {
		r.builder.WriteString(s.String())
	}
}

func (r *Builder) AddString(s lipgloss.Style) {
	if r.params.CurrentLineIsInView() {
		if !r.params.SkipCursorForCurrentLine && r.params.CurrentLine == r.params.Cursor {
			r.builder.WriteString(r.params.Theme.Cursor(s.Value()).String())
		} else {
			r.builder.WriteString(s.String())
		}
	}
}

func (r *Builder) WriteString(s string) {
	if r.params.CurrentLineIsInView() {
		r.builder.WriteString(s)
	}
}

func (r *Builder) writeIndent() {
	for range r.params.Indentation {
		r.builder.WriteRune(' ')
	}
}

func (r *Builder) InsertNewLine() {
	if r.params.CurrentLineIsInView() {

		// this if prevents from rendering the last \n which can break the rendering
		// see https://github.com/charmbracelet/bubbletea/issues/1004
		if r.params.CurrentLine < r.params.ScreenStart+r.params.ScreenHeight-1 {
			r.builder.WriteRune('\n')
		}
	}
	if r.params.NextLineIsInView() {
		r.writeIndent()
	}
}

func (r *Builder) String() string {
	return r.builder.String()
}
