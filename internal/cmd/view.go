package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m *TfxModel) previewLine() string {
	selectionName := m.cursor.Address()

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

func (m *TfxModel) View() string {
	switch m.state {
	case showHelp:
		return m.helpScreen()
	case viewState:
		return m.viewState()
	}
	panic("unknown state")
}

func (m *TfxModel) viewState() string {
	if m.screenHeight == 0 {
		return ""
	}

	var screen []string

	currentNode := m.screenStart
	for i := 0; i < m.screenHeight-1 && currentNode != nil; i++ {
		line := ""
		if currentNode == m.cursor {
			line = m.theme.RenderWithCursor(currentNode.View())
		} else {
			line = m.theme.Render(currentNode.View())
		}
		screen = append(screen, line)

		currentNode = currentNode.Next()
	}

	view := strings.Join(screen, "\n")

	// make the program take up the entire screen if the state is too short
	view = lipgloss.PlaceVertical(m.screenHeight-1, lipgloss.Top, view)

	return fmt.Sprintf("%s\n%s", view, m.previewLine())
}
