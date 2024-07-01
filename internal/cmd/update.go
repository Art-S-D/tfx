package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
)

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
	if previous != m.root {
		m.cursor = previous
	}
	m.moveScreenToCursor()
}

func (m *TfxModel) cursorDown() {
	next := m.cursor.Next()
	if next != nil {
		m.cursor = next
	}
	m.moveScreenToCursor()
}

func (m *TfxModel) goToBottom() {
	m.cursor = m.root.LastChild()
	m.screenStart = m.cursor
	for i := 0; i < m.screenHeight-2; i++ {
		m.screenStart = m.screenStart.Previous()
	}
}

func (m *TfxModel) pageDown() {
	m.cursor = m.screenEnd()
	for i := 0; i < m.screenHeight-1; i++ {
		m.cursorDown()
	}
}
func (m *TfxModel) pageUp() {
	m.cursor = m.screenStart
	for i := 0; i < m.screenHeight-1; i++ {
		m.cursorUp()
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
		case "down":
			m.cursorDown()
		case "g", "G":
			m.goToBottom()
		case "L", "shift+right":
			m.moveScreenToCursor()
			m.cursor.ExpandRecursively()
		case "H", "shift+left":
			m.moveScreenToCursor() // prevents being stuck in the void below the state if a bug section is collapsed
			m.cursor.CollapseRecursively()
			m.clampScreenStart()
		case "?":
			m.state = showHelp
		case "enter":
			m.moveScreenToCursor()
			m.cursor.Expand()
		case "backspace":
			m.moveScreenToCursor()
			m.cursor.Collapse()

			// need to clamp the screen start in case an element at the bottom of the screen collapsed
			m.clampScreenStart()
		case "pgdown", "space", "f":
			m.pageDown()
		case "pgup", "b":
			m.pageUp()
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
