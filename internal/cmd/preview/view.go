package preview

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m *PreviewModel) previewLine() string {
	selectionName := m.cursor.Address()

	helpText := "[?]help [q]quit "
	var previewLine string
	if m.Ctx.ScreenWidth > len(selectionName)+len(helpText) {
		previewLine = fmt.Sprintf(
			"%s%s%s",
			selectionName,
			strings.Repeat(" ", m.Ctx.ScreenWidth-len(selectionName)-len(helpText)),
			helpText,
		)
	} else {
		previewLine = helpText
	}
	return m.Ctx.Theme.Selection(previewLine).String()
}

func (m *PreviewModel) View() string {
	if m.Ctx.ScreenHeight == 0 {
		return ""
	}

	var screen []string

	currentNode := m.screenStart
	for i := 0; i < m.Ctx.ScreenHeight-1 && currentNode != nil; i++ {
		if currentNode.IsRootModule() {
			// skip root module
			currentNode = currentNode.Next()
			continue
		}

		line := ""
		if currentNode == m.cursor {
			line = m.Ctx.Theme.RenderWithCursor(currentNode.View())
		} else {
			line = m.Ctx.Theme.Render(currentNode.View())
		}
		screen = append(screen, line)

		currentNode = currentNode.Next()
	}

	view := strings.Join(screen, "\n")

	// make the program take up the entire screen if the state is too short
	view = lipgloss.PlaceVertical(m.Ctx.ScreenHeight-1, lipgloss.Top, view)

	return fmt.Sprintf("%s\n%s", view, m.previewLine())
}
