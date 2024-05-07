package render

import (
	"fmt"
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

	builder     *strings.Builder
	previewLine string
}

func NewRenderer(cursor, screenStart, screenWidth, screenHeight int, previewLine string) *Renderer {
	var builder strings.Builder
	// fixes a bug where the bottom half of the screen will not rerender when scrolling up
	builder.WriteRune('\n')
	return &Renderer{
		currentLine:              0,
		cursor:                   cursor,
		screenStart:              screenStart,
		screenWidth:              screenWidth,
		screenHeight:             screenHeight,
		currentIndentation:       0,
		builder:                  &builder,
		previewLine:              previewLine,
		skipCursorForCurrentLine: false,
	}
}

func (r *Renderer) currentLineIsInView() bool {
	// the last -1 is to account fot the preview line
	return r.currentLine >= r.screenStart && r.currentLine < r.screenStart+r.screenHeight-1
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
			r.Write(" ")
		}
	}
}
func (r *Renderer) NewLine() {
	r.skipCursorForCurrentLine = false
	r.currentLine += 1

	// this line break is after currentLine+=1 so that the line break is not written on the last lien of the screen
	r.Write("\n")

	r.writeIndent()
}

// sets an internal flat that prevents rendering the cursor on the rest of the current line.
// useful if you want so only show keys and not values as selected
func (r *Renderer) EndCursorForCurrentLine() {
	r.skipCursorForCurrentLine = true
}

func (r *Renderer) String() string {
	// do not write the previewLine to r.builder
	// it will mess up the result if String() is called multiple times
	return fmt.Sprintf("%s\n%s", r.builder.String(), r.previewLine)
}

func (r *Renderer) IndentRight() {
	r.currentIndentation += INDENT_WIDTH
}
func (r *Renderer) IndentLeft() {
	r.currentIndentation -= INDENT_WIDTH
}
