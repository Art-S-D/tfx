package preview

import (
	"github.com/Art-S-D/tfx/internal/cmd/help"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *PreviewModel) screenUp() (hasMoved bool) {
	if m.screenStart == m.node {
		return false
	}
	previous := m.screenStart.Previous()
	m.screenStart = previous
	return true
}
func (m *PreviewModel) screenDown() (hasMoved bool) {
	next := m.screenEnd().Next()
	if next == nil {
		return false
	}
	m.screenStart = m.screenStart.Next()
	return true
}

func (m *PreviewModel) cursorUp() {
	if m.cursor == m.node {
		return
	}
	m.cursor = m.cursor.Previous()
	m.moveScreenToCursor()
}

func (m *PreviewModel) cursorDown() {
	next := m.cursor.Next()
	if next != nil {
		m.cursor = next
	}
	m.moveScreenToCursor()
}

func (m *PreviewModel) goToBottom() {
	m.cursor = m.node.LastChild()
	m.screenStart = m.cursor
	for i := 0; i < m.Ctx.ScreenHeight-2; i++ {
		if m.screenStart != m.node {
			m.screenStart = m.screenStart.Previous()
		}
	}
}

func (m *PreviewModel) pageDown() {
	for i := 0; i < m.Ctx.ScreenHeight-1; i++ {
		m.cursorDown()
	}
}
func (m *PreviewModel) pageUp() {
	for i := 0; i < m.Ctx.ScreenHeight-1; i++ {
		m.cursorUp()
	}
}

func (m *PreviewModel) halfPageDown() {
	for i := 0; i < (m.Ctx.ScreenHeight-1)/2; i++ {
		m.cursorDown()
	}
}
func (m *PreviewModel) halfPageUp() {
	for i := 0; i < (m.Ctx.ScreenHeight-1)/2; i++ {
		m.cursorUp()
	}
}

func (m *PreviewModel) expandAll() {
	m.node.ExpandRecursively()
}
func (m *PreviewModel) collapseAll() {
	m.node.CollapseRecursively()
	for !m.node.HasChild(m.cursor) {
		m.cursor = m.cursor.Parent()
	}
}

func (m *PreviewModel) nextSibling() {
	nextSibling := m.cursor.NextSibling()
	if nextSibling == nil {
		m.cursorDown()
		return
	} else {
		m.cursor = nextSibling
		m.moveScreenToCursor()
	}
}
func (m *PreviewModel) previousSibling() {
	if m.cursor == m.node {
		return
	}
	previousSibling := m.cursor.PreviousSibling()
	if previousSibling != nil {
		m.cursor = previousSibling
		m.moveScreenToCursor()
	}
}

func (m *PreviewModel) printValue() {
	m.cursor.ExpandRecursively()
	m.cursor.IncreaseDepthBy(-m.cursor.Depth())
	m.Ctx.PrintOnExit = m.cursor
}

func (m *PreviewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.Ctx.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q", "esc":
			if m.ParentModel == nil {
				return m, tea.Quit
			}
			return m.ParentModel, nil
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
			return &help.HelpModel{
				Ctx:         m.Ctx,
				ParentModel: m,
			}, nil
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
		case "p":
			nextModel := &PreviewModel{
				Ctx:         m.Ctx,
				ParentModel: m,
			}
			node := m.cursor.Clone()
			node.Expand()
			node.SetKey("", 0)
			node.IncreaseDepthBy(-node.Depth())
			nextModel.SetNode(node)
			return nextModel, nil
		case "r":
			m.cursor.Reveal()
		}
	case tea.WindowSizeMsg:
		m.clampScreenStart()
	case tea.MouseMsg:
		m.handleMouseEvent(msg)
	}
	return m, nil
}
