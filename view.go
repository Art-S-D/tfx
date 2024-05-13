package main

import (
	"fmt"
	"strings"

	"github.com/Art-S-D/tfx/internal/render"
	"github.com/charmbracelet/lipgloss"
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
	return m.theme.Selection(previewLine).String()
}

func (m *stateModel) View() string {
	preview := m.previewLine()
	params := render.ViewParams{
		CurrentLine:              0,
		Indentation:              0,
		SkipCursorForCurrentLine: false,

		Theme:        m.theme,
		Cursor:       m.cursor,
		ScreenStart:  m.screenStart,
		ScreenWidth:  m.screenWidth,
		ScreenHeight: m.screenHeight - 1,
	}
	screen := m.rootModule.View(&params)
	screen = lipgloss.PlaceVertical(m.screenHeight-1, lipgloss.Top, screen)

	return fmt.Sprintf("%s\n%s", screen, preview)
}
