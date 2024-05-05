package render

import (
	"strings"

	"github.com/Art-S-D/tfview/internal/style"
)

const INDENT_WIDTH = 2

type ViewParams struct {
	Cursor int
	Width  int
	// Height int
	Indentation int
}

func (v ViewParams) Indent() string {
	return strings.Repeat(" ", v.Indentation)
}
func (v ViewParams) NextIndent() ViewParams {
	out := v
	out.Indentation += INDENT_WIDTH
	return out
}
func (v ViewParams) PrevIndent() ViewParams {
	out := v
	out.Indentation -= INDENT_WIDTH
	return out
}

type Renderer struct {
	currentLine int
	cursor      int

	screenStart  int // should always be between [0, rootModule.Height() - screenHeight)
	screenWidth  int
	screenHeight int

	currentIndentation int

	builder *strings.Builder
}

func (r *Renderer) currentLineIsInView() bool {
	return r.currentLine >= r.screenStart && r.currentLine > r.screenStart+r.screenHeight
}
func (r *Renderer) Write(s string) {
	if r.currentLineIsInView() {
		r.builder.WriteString(s)
	}
}

// applies the cursor style if the current line is selected
func (r *Renderer) WriteCursor(s string) {
	if r.cursor == r.currentLine {
		r.Write(style.Cursor.Render(s))
	} else {
		r.Write(s)
	}
}

func (r *Renderer) writeIndent() {
	if r.currentLineIsInView() {
		for range r.currentIndentation {
			r.Write(" ")
		}
	}
}
func (r *Renderer) NewLine() {
	r.Write("\n")
	r.currentLine += 1
	r.writeIndent()

}
func (r *Renderer) String() string {
	return r.builder.String()
}

func (r *Renderer) IndentRight() {
	r.currentIndentation += INDENT_WIDTH
}
func (r *Renderer) IndentLeft() {
	r.currentIndentation -= INDENT_WIDTH
}

type Model interface {
	View(opts ViewParams) string
	Selected(cursor int) (selected Model, cursorPosition int)
	Address() string
	Expand()
	Collapse()
	ViewHeight() int
}
