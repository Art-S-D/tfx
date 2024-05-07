package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

// func (m *stateModel) screenBottom() int {
// 	// should be -1 but we have -2 instead to account for the preview
// 	return m.screenHeight + m.screenStart - 2
// }

// moves the screen so that the cursor is in view
// used when the ui jumps around, ex: when a big element is collapsed
// basically the cursor does what it wants and the screen follows it
func (m *stateModel) clampScreen() {
	minScreen := 0
	maxScreen := m.rootModuleHeight - m.screenHeight + 1
	if m.screenStart > m.cursor {
		m.screenStart = max(minScreen, m.cursor)
	}

	// +2 because we need one for the preview line and one so that the cursor is one line on top of the bottom of the screen
	if m.screenStart < m.cursor-m.screenHeight+2 {
		m.screenStart = min(maxScreen, m.cursor-m.screenHeight+2)
	}
}

// moves the cursor so that it does not go out of the state
func (m *stateModel) clampCursor() {
	if m.cursor < 0 {
		m.cursor = 0
	}
	if m.cursor > m.rootModuleHeight-1 {
		m.cursor = m.rootModuleHeight - 1
	}
}

func (m *stateModel) cursorUp() {
	m.cursor -= 1
	m.clampCursor()
	m.clampScreen()
}

func (m *stateModel) cursorDown() {
	m.cursor += 1
	m.clampCursor()
	m.clampScreen()
}

func (m *stateModel) goToBottom() {
	m.cursor = m.rootModuleHeight - 1
}

// func (m *stateModel) pageDown() {
// 	m.screenStart += m.screenHeight
// 	m.cursor = m.screenStart
// 	m.clampScreenStart()
// 	m.clampCursor()
// }

func (m *stateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.clampScreen()
		case "enter":
			selected, _ := m.rootModule.Selected(m.cursor)
			selected.Expand()
			m.rootModuleHeight = m.rootModule.ViewHeight()
			m.clampCursor()
			m.clampScreen()
		case "backspace":
			selected, selectedLine := m.rootModule.Selected(m.cursor)
			selected.Collapse()
			m.cursor -= selectedLine
			m.rootModuleHeight = m.rootModule.ViewHeight()
			m.clampCursor()
			m.clampScreen()
		}
	case tea.WindowSizeMsg:
		m.screenHeight = msg.Height
		m.screenWidth = msg.Width

		m.clampCursor()
		m.clampScreen()
	}
	return m, nil
}
