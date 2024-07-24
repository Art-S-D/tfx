package preview

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *PreviewModel) flipExpandedAtCursor() {
	if m.cursor.IsExpanded() {
		m.cursor.Collapse()
		m.clampScreenStart()
	} else {
		m.cursor.Expand()
	}
}

func (m *PreviewModel) handleMouseEvent(msg tea.MouseMsg) {
	if msg.Action != tea.MouseActionPress {
		return
	}
	switch msg.Button {
	case tea.MouseButtonWheelDown:
		m.screenDown()
	case tea.MouseButtonWheelUp:
		m.screenUp()
	case tea.MouseButtonLeft:
		// prevent clicking on the preview line
		if msg.Y >= m.Ctx.ScreenHeight-1 {
			return
		}

		destination := m.screenStart
		if m.screenStart.IsRootModule() {
			destination = destination.Next()
		}
		for i := 0; i < msg.Y && destination != nil; i++ {
			destination = destination.Next()
		}
		if destination == nil {
			// clicked on an empty line below the screen
			return
		}

		if m.cursor == destination {
			// if the row is already selected, expand/collapse it
			m.flipExpandedAtCursor()
		} else {
			m.cursor = destination
		}
	}
}
