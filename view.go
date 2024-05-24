package main

import (
	"fmt"
	"strings"

	"github.com/Art-S-D/tfx/internal/render"
)

func (m *stateModel) Selected() render.Model {
	currentLine := m.screen[m.cursor]
	return currentLine.PointsTo
}

func (m *stateModel) previewLine() string {
	selection := m.Selected()
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
	return m.theme.Selection(previewLine).String()
}

func (m *stateModel) View() string {
	switch m.state {
	case showHelp:
		return m.helpScreen()
	case viewState:
		return m.viewState()
	}
	panic("unknown state")
}

func (m *stateModel) viewState() string {
	if m.screenHeight == 0 {
		return ""
	}

	screen := m.rootModule.View()
	screenSlice := screen[m.screenStart : m.screenStart+m.screenHeight-1]
	var sb strings.Builder
	for i, line := range screenSlice {
		sb.WriteString(line.Render(m.theme, i+m.screenStart == m.cursor))
		sb.WriteRune('\n')
	}
	sb.WriteString(m.previewLine())
	// screen := lipgloss.PlaceVertical(m.screenHeight-1, lipgloss.Top, screen)

	return sb.String()
}
