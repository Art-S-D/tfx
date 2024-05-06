package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *stateModel) screenBottom() int {
	// should be -1 but we have -2 instead to account for the preview
	return m.screenHeight + m.screenStart - 2
}

func (m *stateModel) clampScreenStart() {
	if m.screenStart < 0 {
		m.screenStart = 0
	}
	if m.screenStart >= m.rootModuleHeight-m.screenHeight {
		m.screenStart = m.rootModuleHeight - m.screenHeight
	}
}

func (m *stateModel) clampCursor() {
	if m.cursor < 0 {
		m.cursor = 0
	}
	screenBottom := m.screenBottom()
	if m.cursor >= screenBottom {
		m.cursor = screenBottom
	}
}

func (m *stateModel) cursorUp() {
	m.cursor -= 1
	if m.cursor < 0 {
		m.cursor = 0
	}
	if m.cursor <= m.screenStart+m.screenDrag {
		m.screenStart -= 1
	}
	m.clampScreenStart()
	m.clampCursor()
}

func (m *stateModel) cursorDown() {
	m.cursor += 1
	screenBottom := m.screenBottom()
	if m.cursor >= screenBottom {
		m.cursor = screenBottom
	}
	if m.cursor >= screenBottom-m.screenDrag {
		m.screenStart += 1
	}
	if m.cursor < m.screenStart+m.screenDrag {
		m.screenStart -= 1
	}
	m.clampScreenStart()
	m.clampCursor()
}

func (m *stateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up":
			m.cursorUp()
			return m, nil
		case "down":
			m.cursorDown()
			return m, nil
		case "enter":
			selected, _ := m.rootModule.Selected(m.cursor)
			selected.Expand()
			m.rootModuleHeight = m.rootModule.ViewHeight()
			m.clampScreenStart()
			m.clampCursor()
		case "backspace":
			selected, selectedLine := m.rootModule.Selected(m.cursor)
			selected.Collapse()
			m.cursor -= selectedLine
			m.rootModuleHeight = m.rootModule.ViewHeight()
			m.clampScreenStart()
			m.clampCursor()
		}
	case tea.WindowSizeMsg:
		m.screenHeight = msg.Height
		m.screenWidth = msg.Width

		m.clampScreenStart()
		m.clampCursor()
	}
	return m, nil
}
