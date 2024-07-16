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
	for i := 0; i < m.screenHeight-1; i++ {
		m.cursorDown()
	}
}
func (m *TfxModel) pageUp() {
	for i := 0; i < m.screenHeight-1; i++ {
		m.cursorUp()
	}
}

func (m *TfxModel) halfPageDown() {
	for i := 0; i < (m.screenHeight-1)/2; i++ {
		m.cursorDown()
	}
}
func (m *TfxModel) halfPageUp() {
	for i := 0; i < (m.screenHeight-1)/2; i++ {
		m.cursorUp()
	}
}

func (m *TfxModel) expandAll() {
	m.root.ExpandRecursively()
}
func (m *TfxModel) collapseAll() {
	m.root.CollapseRecursively()
	for !m.root.HasChild(m.cursor) {
		m.cursor = m.cursor.Parent
	}
}

func (m *TfxModel) nextSibling() {
	nextSibling := m.cursor.NextSibling()
	if nextSibling == nil {
		m.cursorDown()
		return
	} else {
		m.cursor = nextSibling
		m.moveScreenToCursor()
	}
}
func (m *TfxModel) previousSibling() {
	previousSibling := m.cursor.PreviousSibling()
	if previousSibling != nil && previousSibling != m.root {
		m.cursor = previousSibling
		m.moveScreenToCursor()
	}
}

func (m *TfxModel) printValue() {
	m.cursor.ExpandRecursively()
	m.cursor.IndentBy(-m.cursor.Depth)
	m.PrintOnExit = m.cursor
}

func (m *TfxModel) updateStateView(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up", "k":
			m.cursorUp()
		case "down", "j":
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
		case "enter", "l", "right":
			m.moveScreenToCursor()
			m.cursor.Expand()
		case "backspace", "h", "left":
			m.moveScreenToCursor()
			m.cursor.Collapse()
			// need to clamp the screen start in case an element at the bottom of the screen collapsed
			m.clampScreenStart()
		case "pgdown", "space", "f":
			m.pageDown()
		case "pgup", "b":
			m.pageUp()
		case "u", "ctrl+u":
			m.halfPageUp()
		case "d", "ctrl+d":
			m.halfPageDown()
		case "E":
			m.collapseAll()
			m.moveScreenToCursor()
		case "e":
			m.expandAll()
			m.moveScreenToCursor()
		case "J", "shift+down":
			m.nextSibling()
		case "K", "shift+up":
			m.previousSibling()
		case "P":
			m.printValue()
			return m, tea.Quit
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
