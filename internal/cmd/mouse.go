package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *TfxModel) flipExpandedAtCursor() {
	if m.cursor.IsExpanded() {
		m.cursor.Collapse()
		m.clampScreenStart()
	} else {
		m.cursor.Expand()
	}
}

func (m *TfxModel) moveScreenToCursor() {
	end := m.screenStart
	for i := 0; i < m.screenHeight-2 && end.Next() != nil; i++ {
		if end == m.cursor {
			// the cursor is already in view => stop here
			return
		}
		end = end.Next()
	}

	// try to scroll down to the cursor
	current := end
	start := m.screenStart
	for current != nil {
		if current == m.cursor {
			m.screenStart = start
			return
		}
		current = current.Next()
		start = start.Next()
	}

	// try to scroll up to the cursor
	current = m.screenStart
	for current != m.root {
		if current == m.cursor {
			m.screenStart = current
			return
		}
		current = current.Previous()
	}

	panic("cursor not found")
}

func (m *TfxModel) handleMouseEvent(msg tea.MouseMsg) {
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
		if msg.Y >= m.screenHeight-1 {
			return
		}

		destination := m.screenStart
		for i := 0; i < msg.Y; i++ {
			destination = destination.Next()
		}
		if m.cursor == destination {
			// if the row is already selected, expand/collapse it
			m.flipExpandedAtCursor()
		} else {
			m.cursor = destination
		}
	}
}
