package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
)

//	func (m *TfxModel) screenBottom() int {
//		// should be -1 but we have -2 instead to account for the preview
//		return m.screenHeight + m.screenStart - 2
//	}
func (m *TfxModel) clampScreenStart() {
	end := m.screenStart
	for i := 0; i < m.screenHeight-2; i++ {
		next := end.Next()
		if next == nil {
			// screenEnd is at the end of the screen => we need to move the screenStart up
			previous := m.screenStart.Previous()
			if previous == m.root {
				// screenEnd is at the end and screenStart is at the start
				// => meaning the state is too small to cover the entire screen
				// so we can just stop
				return
			}
			m.screenStart = previous
		} else {
			// otherwise just go down
			end = end.Next()
		}
	}
}

func (m *TfxModel) screenUp() (hasMoved bool) {
	previous := m.screenStart.Previous()
	if previous == m.root {
		return false
	}
	m.screenStart = previous
	return true
}
func (m *TfxModel) screenDown() (hasMoved bool) {
	next := m.screenEnd().Next()
	if next == nil {
		return false
	}
	m.screenStart = m.screenStart.Next()
	return true
}

func (m *TfxModel) cursorUp() {
	previous := m.cursor.Previous()
	if previous == m.root {
		return
	}
	if m.cursor == m.screenStart {
		// move the screen with the cursor
		m.cursor = previous
		m.screenStart = previous
	} else {
		m.cursor = previous
	}
}

func (m *TfxModel) cursorDown() {
	next := m.cursor.Next()
	if next == nil {
		return
	}
	if m.cursor == m.screenEnd() {
		// move the screen with the cursor
		m.cursor = next
		m.screenStart = m.screenStart.Next()
	} else {
		m.cursor = next
	}
}

func (m *TfxModel) goToBottom() {
	m.cursor = m.root.LastChild()
	m.screenStart = m.cursor
	for i := 0; i < m.screenHeight-2; i++ {
		m.screenStart = m.screenStart.Previous()
	}
}

func (m *TfxModel) updateStateView(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up":
			m.cursorUp()
			m.moveScreenToCursor()
		case "down":
			m.cursorDown()
			m.moveScreenToCursor()
		case "g", "G":
			m.goToBottom()
		case "L", "shift+right":
			m.cursor.ExpandRecursively()
		case "H", "shift+left":
			m.cursor.CollapseRecursively()
		case "?":
			m.state = showHelp
		case "enter":
			m.cursor.Expand()
		case "backspace":
			m.cursor.Collapse()

			// need to clamp the screen start in case an element at the bottom of the screen collapsed
			m.clampScreenStart()
		case "r":
			m.cursor.Reveal()
		}
	case tea.WindowSizeMsg:
		m.screenHeight = msg.Height
		m.screenWidth = msg.Width
		m.clampScreenStart()
	case tea.MouseMsg:
		m.handleMouseEvent(msg)
	}
	return m, nil
}

func (m *TfxModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case showHelp:
		return m.updateHelpView(msg)
	case viewState:
		return m.updateStateView(msg)
	}
	panic("unknown state")
}
