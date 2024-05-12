package main

import (
	"fmt"
	"strings"

	"github.com/Art-S-D/tfx/internal/style"
	"github.com/charmbracelet/lipgloss"
)

func (m *stateModel) previewLine() string {
	selection := m.selected()
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

	// this happens before the first update message is sent to the model
	if m.screenHeight == 0 {
		return ""
	}

	preview := m.previewLine()

	renderedLines := m.screen[m.screenStart:min(len(m.screen), m.screenStart+m.screenHeight-1)]
	var sb strings.Builder
	for i, line := range renderedLines {
		lineNumber := m.screenStart + i
		sb.WriteString(line.View(lineNumber == m.cursor))
		if i < len(renderedLines)-1 {
			sb.WriteRune('\n')
		}
	}
	screen := sb.String()
	screen = lipgloss.PlaceVertical(m.screenHeight-1, lipgloss.Top, screen)

	return fmt.Sprintf("%s\n%s", screen, preview)
}
