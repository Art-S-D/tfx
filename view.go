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
	renderer := render.NewRenderer(
		m.cursor,
		m.screenStart,
		m.screenWidth,
		m.screenHeight,
		m.previewLine(),
	)
	m.rootModule.View(renderer)
	return renderer.String()
}
