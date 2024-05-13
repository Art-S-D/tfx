package render

import "github.com/Art-S-D/tfx/internal/style"

const INDENT_WIDTH = 2

// a wrapper around strings.Builder that can manage indentation and rendering the Cursor
type ViewParams struct {
	CurrentLine int
	Indentation int

	Cursor                   int
	SkipCursorForCurrentLine bool

	Theme *style.Theme

	ScreenStart  int
	ScreenWidth  int
	ScreenHeight int
}

func NewViewParams(Cursor, ScreenStart, ScreenWidth, ScreenHeight int) *ViewParams {
	return &ViewParams{
		CurrentLine:              0,
		Cursor:                   Cursor,
		ScreenStart:              ScreenStart,
		ScreenWidth:              ScreenWidth,
		ScreenHeight:             ScreenHeight,
		Indentation:              0,
		SkipCursorForCurrentLine: false,
	}
}

func (r *ViewParams) CurrentLineIsInView() bool {
	return r.CurrentLine >= r.ScreenStart && r.CurrentLine < r.ScreenStart+r.ScreenHeight
}
func (r *ViewParams) NextLineIsInView() bool {
	line := r.CurrentLine + 1
	return line >= r.ScreenStart && line < r.ScreenStart+r.ScreenHeight
}

// advance the line number
func (r *ViewParams) NextLine() {
	r.SkipCursorForCurrentLine = false
	r.CurrentLine += 1
}

// sets an internal flat that prevents rendering the Cursor on the rest of the current line.
// useful if you want so only show keys and not values as selected
func (r *ViewParams) EndCursorForCurrentLine() {
	r.SkipCursorForCurrentLine = true
}

// clones r
func (r *ViewParams) IndentRight() {
	r.Indentation += INDENT_WIDTH
	r.ScreenWidth -= INDENT_WIDTH
}

// clones r
func (r *ViewParams) IndentLeft() {
	r.Indentation -= INDENT_WIDTH
	r.ScreenWidth += INDENT_WIDTH
}
