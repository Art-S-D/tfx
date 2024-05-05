package main

import tea "github.com/charmbracelet/bubbletea"

func (m *stateModel) cursorUp() {
	m.cursor -= 1
	if m.cursor < 0 {
		m.cursor = 0
	}
	if m.cursor <= m.offset+m.screenDrag {
		m.offset -= 1
	}
	m.clampOffset()
	m.clampCursor()
}

func (m *stateModel) cursorDown() {
	m.cursor += 1
	screenBottom := m.screenBottom()
	if m.cursor >= screenBottom {
		m.cursor = screenBottom
	}
	if m.cursor >= screenBottom-m.screenDrag {
		m.offset += 1
	}
	m.clampOffset()
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
			m.clampOffset()
			m.clampCursor()
		case "backspace":
			selected, selectedLine := m.rootModule.Selected(m.cursor)
			selected.Collapse()
			m.cursor -= selectedLine
			m.rootModuleHeight = m.rootModule.ViewHeight()
			m.clampOffset()
			m.clampCursor()
		}
	case tea.WindowSizeMsg:
		m.screenHeight = msg.Height
		m.screenWidth = msg.Width

		m.clampOffset()
	}
	return m, nil
}
