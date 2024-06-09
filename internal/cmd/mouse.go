package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *TfxModel) flipExpandedAtCursor() {
	selected := m.Selected()
	if selected.IsCollapsed() {
		m.expandAtSelection()
	} else {
		m.collapseAtSelection()
	}
}
func (m *TfxModel) handleMouseEvent(msg tea.MouseMsg) {
	if msg.Action != tea.MouseActionPress {
		return
	}
	switch msg.Button {
	case tea.MouseButtonWheelDown:
		m.screenStart += 1
		m.clampScreen()
	case tea.MouseButtonWheelUp:
		m.screenStart -= 1
		m.clampScreen()
	case tea.MouseButtonLeft:
		destinationRow := m.screenStart + msg.Y
		if destinationRow > m.EntireHeight() {
			return
		}

		if m.cursor == destinationRow {
			// if the row is already selected, expand/collapse it
			m.flipExpandedAtCursor()
		} else {
			m.cursor = destinationRow
		}
	}
}
