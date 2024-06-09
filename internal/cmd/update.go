package cmd

import (
	"slices"

	"github.com/Art-S-D/tfx/internal/json"
	"github.com/Art-S-D/tfx/internal/render"
	tea "github.com/charmbracelet/bubbletea"
)

// func (m *TfxModel) screenBottom() int {
// 	// should be -1 but we have -2 instead to account for the preview
// 	return m.screenHeight + m.screenStart - 2
// }

func (m *TfxModel) EntireHeight() int {
	return len(m.screen)
}

func (m *TfxModel) clampScreen() {
	// can be negative but this will be fixed by max(0, m.screenStart)
	// this is important when collapsing a node at the bottom of the screen
	lowestPossibleScreen := m.EntireHeight() - m.screenHeight + 1
	m.screenStart = min(lowestPossibleScreen, m.screenStart)
	m.screenStart = max(0, m.screenStart)
}

// moves the cursor so that it does not go out of the state
func (m *TfxModel) clampCursor() {
	if m.cursor < 0 {
		m.cursor = 0
	}
	if m.cursor > m.EntireHeight()-1 {
		m.cursor = m.EntireHeight() - 1
	}
}

// moves the screen so that the cursor is in view
// used when the ui jumps around, ex: when a big element is collapsed
// basically the cursor does what it wants and the screen follows it
func (m *TfxModel) moveScreenToCursor() {
	m.clampScreen()

	// cursor is above the screen
	if m.cursor < m.screenStart {
		m.screenStart = max(0, m.cursor)
	}

	// cursor is below the screen
	// -2 because we need one for the preview line and one so that the cursor is one line on top of the bottom of the screen
	screenBottom := m.screenStart + m.screenHeight - 2
	maxScreen := max(0, m.EntireHeight()-m.screenHeight+1)
	if m.cursor >= screenBottom {
		m.screenStart = min(maxScreen, m.cursor-m.screenHeight+2)
	}
}

func (m *TfxModel) cursorUp() {
	m.cursor -= 1
	m.clampCursor()
	m.moveScreenToCursor()
}

func (m *TfxModel) cursorDown() {
	m.cursor += 1
	m.clampCursor()
	m.moveScreenToCursor()
}

func (m *TfxModel) goToBottom() {
	m.cursor = m.EntireHeight() - 1
	m.moveScreenToCursor()
}

// func (m *TfxModel) pageDown() {
// 	m.screenStart += m.screenHeight
// 	m.cursor = m.screenStart
// 	m.moveScreenToCursorStart()
// 	m.clampCursor()
// }

func (m *TfxModel) revealSensitiveValue(model render.Model, value *json.SensitiveValue) {
	height := len(value.View())
	value.Reveal()
	m.replaceAtCursor(model, height)

	m.clampCursor()
	m.moveScreenToCursor()
}

// replace the `previousHeight`th lines under the cursor by the View of `by`
// while keeping the previous indentation
// this allows swapping the previous content of a Model by the new content
func (m *TfxModel) replaceAtCursor(by render.Model, previousHeight int) {
	indentation := m.screen[m.cursor].Indentation
	nextLines := by.View()
	render.IndentBy(nextLines, indentation)
	m.screen = slices.Replace(m.screen, m.cursor, m.cursor+previousHeight, nextLines...)
}

func (m *TfxModel) expandAtSelection() {
	selected := m.Selected()
	previousHeight := len(selected.View())
	selected.Expand()
	m.replaceAtCursor(selected, previousHeight)

	m.clampCursor()
	m.moveScreenToCursor()

}
func (m *TfxModel) collapseAtSelection() {
	selected := m.Selected()
	previousLines := selected.View()
	if m.screen[m.cursor].PointsToEnd {
		m.cursor -= len(previousLines) - 1
	}
	selected.Collapse()
	m.replaceAtCursor(selected, len(previousLines))

	m.clampCursor()
	m.moveScreenToCursor()
}
func (m *TfxModel) revealAtSelection() {
	selected := m.Selected()
	if sensitiveValue, ok := selected.(*json.SensitiveValue); ok {
		m.revealSensitiveValue(selected, sensitiveValue)
	}
	if kv, ok := selected.(*json.KeyVal); ok {
		if sensitiveValue, ok := kv.Value.(*json.SensitiveValue); ok {
			m.revealSensitiveValue(selected, sensitiveValue)
		}
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
		case "?":
			m.state = showHelp
		case "enter":
			m.expandAtSelection()
		case "backspace":
			m.collapseAtSelection()
		case "r":
			m.revealAtSelection()
		}
	case tea.WindowSizeMsg:
		m.screenHeight = msg.Height
		m.screenWidth = msg.Width

		m.clampCursor()
		m.moveScreenToCursor()
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
