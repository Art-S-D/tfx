package main

import (
	"fmt"
	"strings"

	"github.com/Art-S-D/tfview/internal/render"
	"github.com/Art-S-D/tfview/internal/style"
)

func (m *stateModel) View() string {
	view := m.rootModule.View(render.ViewParams{Cursor: m.cursor, Width: m.screenWidth, Indentation: 0})
	lines := strings.Split(view, "\n")
	linesInView := lines[m.offset : m.offset+m.screenHeight]
	if len(linesInView) > 1 {
		selection, _ := m.rootModule.Selected(m.cursor)
		selectionName := selection.Address()

		helpText := "[?]help [q]quit "
		previewLine := fmt.Sprintf(
			"%s%s%s",
			selectionName,
			strings.Repeat(" ", m.screenWidth-len(selectionName)-len(helpText)),
			helpText,
		)

		linesInView[len(linesInView)-1] = style.Selection.Copy().
			PaddingRight(m.screenWidth - len(selectionName)).
			Render(previewLine)
	}
	return strings.Join(linesInView, "\n")
}
