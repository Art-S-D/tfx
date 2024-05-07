package main

import (
	"fmt"
	"strings"

	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
)

func (m *stateModel) previewLine() string {
	selection, _ := m.rootModule.Selected(m.cursor)
	selectionName := selection.Address()

	helpText := "[?]help [q]quit "
	var previewLine string
	if m.screenWidth > len(selectionName)+len(helpText) {
		previewLine = fmt.Sprintf(
			"%s%s%s",
			selectionName,
			strings.Repeat(" ", m.screenWidth-len(selectionName)-len(helpText)),
			helpText,
		)
	} else {
		previewLine = helpText
	}
	return style.Selection.Render(previewLine)
}

func (m *stateModel) View() string {
	preview := m.previewLine()
	renderer := render.NewRenderer(
		m.cursor,
		m.screenStart,
		m.screenWidth,
		m.screenHeight-1, // -1 to make space for the preview line
	)
	m.rootModule.View(renderer)

	return fmt.Sprintf("%s\n%s", renderer.String(), preview)
}
