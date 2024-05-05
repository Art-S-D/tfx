package render

import (
	"strings"

	"github.com/Art-S-D/tfview/internal/style"
	"github.com/charmbracelet/lipgloss"
)

const INDENT_WIDTH = 2

type Renderer struct {
	currentLine int
	cursor      int

	screenStart  int // should always be between [0, rootModule.Height() - screenHeight)
	screenWidth  int
	screenHeight int

	currentIndentation int

	builder     *strings.Builder
	previewLine string
}

func NewRenderer(cursor, screenStart, screenWidth, screenHeight int, previewLine string) *Renderer {
	return &Renderer{
		currentLine:        0,
		cursor:             cursor,
		screenStart:        screenStart,
		screenWidth:        screenWidth,
		screenHeight:       screenHeight,
		currentIndentation: 0,
		builder:            &strings.Builder{},
		previewLine:        previewLine,
	}
}

func (r *Renderer) currentLineIsInView() bool {
	// the last -1 is to account fot the preview line
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
	if r.cursor == r.currentLine {
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
	r.Write("\n")
	r.currentLine += 1
	r.writeIndent()

}
func (r *Renderer) String() string {
	r.builder.WriteRune('\n')
	r.builder.WriteString(r.previewLine)
	return r.builder.String()
}

func (r *Renderer) IndentRight() {
	r.currentIndentation += INDENT_WIDTH
}
func (r *Renderer) IndentLeft() {
	r.currentIndentation -= INDENT_WIDTH
}

type Model interface {
	View(opts *Renderer)
	Selected(cursor int) (selected Model, cursorPosition int)
	Address() string
	Expand()
	Collapse()
	ViewHeight() int
}
