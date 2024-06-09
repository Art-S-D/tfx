package cmd

import (
	"fmt"
	"strings"

	"github.com/Art-S-D/tfx/internal/render"
	"github.com/charmbracelet/lipgloss"
)

func (m *TfxModel) Selected() render.Model {
	currentLine := m.screen[m.cursor]
	return currentLine.PointsTo
}

func (m *TfxModel) previewLine() string {
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

	screenLines := m.rootModule.View()
	screenBottom := min(m.screenStart+m.screenHeight-1, m.EntireHeight())
	screenSlice := screenLines[m.screenStart:screenBottom]
	var sb strings.Builder
	for i, line := range screenSlice {
		sb.WriteString(line.Render(m.theme, i+m.screenStart == m.cursor))
		if i < len(screenSlice)-1 {
			sb.WriteRune('\n')
		}
	}

	screen := lipgloss.PlaceVertical(m.screenHeight-1, lipgloss.Top, sb.String())
	return fmt.Sprintf("%s\n%s", screen, m.previewLine())
}
