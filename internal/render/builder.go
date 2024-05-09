package render

import (
	"strings"

	"github.com/Art-S-D/tfx/internal/style"
	"github.com/charmbracelet/lipgloss"
)

// a wrapper around strings.Builder that can manage indentation and rendering the cursor
type Builder struct {
	renderingCursor bool
	builder         strings.Builder

	params *ViewParams
}

func NewBuilder(params *ViewParams) *Builder {
	var builder strings.Builder
	return &Builder{
		params:          params,
		builder:         builder,
		renderingCursor: false,
	}
}

func (r *Builder) Reset() {
	r.builder.Reset()
	r.renderingCursor = false
}

func (r *Builder) CursorStart() {
	r.renderingCursor = true
}
func (r *Builder) CursorEnd() {
	r.renderingCursor = true
}

func (r *Builder) WriteString(str string) {
	if r.params.CurrentLineIsInView() {
		if r.renderingCursor {
			r.builder.WriteString(style.Cursor.Render(str))
		} else {
			r.builder.WriteString(str)
		}
	}
}

func (r *Builder) WriteStyle(s lipgloss.Style, str string) {
	if r.params.CurrentLineIsInView() {
		if r.renderingCursor {
			r.builder.WriteString(style.Cursor.Render(str))
		} else {
			r.builder.WriteString(s.Render(str))
		}
	}
}

func (r *Builder) writeIndent() {
	if r.params.CurrentLineIsInView() {
		for range r.params.Indentation {
			r.builder.WriteRune(' ')
		}
	}
}
func (r *Builder) NewLine() {
	if r.params.CurrentLineIsInView() {

		// this if prevents from rendering the last \n which can break the rendering
		// see https://github.com/charmbracelet/bubbletea/issues/1004
		if r.params.CurrentLine < r.params.ScreenStart+r.params.ScreenHeight-1 {
			r.builder.WriteRune('\n')
		}
		r.writeIndent()
	}
}

func (r *Builder) String() string {
	// since components dont end with a new line, we need to insert the last line
	r.NewLine()
	return r.builder.String()
}
