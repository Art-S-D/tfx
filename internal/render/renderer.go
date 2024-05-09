package render

import (
	"strings"

	"github.com/Art-S-D/tfx/internal/style"
	"github.com/charmbracelet/lipgloss"
)

const INDENT_WIDTH = 2

// a wrapper around strings.Builder that can manage indentation and rendering the cursor
type Renderer struct {
	currentLine int
	cursor      int

	screenStart  int
	screenWidth  int
	screenHeight int

	currentIndentation int

	skipCursorForCurrentLine bool

	builder *strings.Builder
}

func NewRenderer(cursor, screenStart, screenWidth, screenHeight int) *Renderer {
	var builder strings.Builder
	// fixes a bug where the bottom half of the screen will not rerender when scrolling up
	// builder.WriteRune('\n')
	return &Renderer{
		currentLine:              0,
		cursor:                   cursor,
		screenStart:              screenStart,
		screenWidth:              screenWidth,
		screenHeight:             screenHeight,
		currentIndentation:       0,
		builder:                  &builder,
		skipCursorForCurrentLine: false,
	}
}

func (r *Renderer) currentLineIsInView() bool {
	return r.currentLine >= r.screenStart && r.currentLine < r.screenStart+r.screenHeight
}

func (r *Renderer) Write(s string) {
	if r.currentLineIsInView() {
		r.builder.WriteString(s)
	}
}

// applies the cursor style if the current line is selected
// or the s style otherwise
func (r *Renderer) CursorWrite(s lipgloss.Style, str string) {
	if r.cursor == r.currentLine && !r.skipCursorForCurrentLine {
		r.Write(style.Cursor.Render(str))
	} else {
		r.Write(s.Render(str))
	}
}

func (r *Renderer) writeIndent() {
	if r.currentLineIsInView() {
		for range r.currentIndentation {
			r.builder.WriteRune(' ')
		}
	}
}
func (r *Renderer) NewLine() {
	if r.currentLineIsInView() {

		// this if prevents from rendering the last \n which can break the rendering
		// see https://github.com/charmbracelet/bubbletea/issues/1004
		if r.currentLine < r.screenStart+r.screenHeight-1 {
			r.builder.WriteRune('\n')
		}
		r.writeIndent()
	}
	r.skipCursorForCurrentLine = false
	r.currentLine += 1
}

// sets an internal flat that prevents rendering the cursor on the rest of the current line.
// useful if you want so only show keys and not values as selected
func (r *Renderer) EndCursorForCurrentLine() {
	r.skipCursorForCurrentLine = true
}

func (r *Renderer) String() string {
	// since components dont end with a new line, we need to insert the last line
	r.NewLine()

	for r.currentLine < r.screenHeight-1 {
		r.currentLine += 1
		r.builder.WriteByte('\n')
	}
	return r.builder.String()
}

func (r *Renderer) IndentRight() {
	r.currentIndentation += INDENT_WIDTH
}
func (r *Renderer) IndentLeft() {
	r.currentIndentation -= INDENT_WIDTH
}
