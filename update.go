package main

import (
	"slices"

	"github.com/Art-S-D/tfx/internal/json"
	"github.com/Art-S-D/tfx/internal/render"
	tea "github.com/charmbracelet/bubbletea"
)

// func (m *stateModel) screenBottom() int {
// 	// should be -1 but we have -2 instead to account for the preview
// 	return m.screenHeight + m.screenStart - 2
// }

func (m *stateModel) ScreenHeight() int {
	return len(m.screen)
}

// moves the screen so that the cursor is in view
// used when the ui jumps around, ex: when a big element is collapsed
// basically the cursor does what it wants and the screen follows it
func (m *stateModel) clampScreen() {
	minScreen := 0
	maxScreen := m.ScreenHeight() - m.screenHeight + 1

	m.screenStart = max(m.screenStart, minScreen)
	m.screenStart = min(m.screenStart, maxScreen) // this is important when collapsing a node at the bottom of the screen

	// cursor is above the screen
	if m.screenStart > m.cursor {
		m.screenStart = max(minScreen, m.cursor)
	}

	// cursor is below the screen
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
	if m.cursor > m.ScreenHeight()-1 {
		m.cursor = m.ScreenHeight() - 1
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
	m.cursor = m.ScreenHeight() - 1
	m.clampScreen()
}

// func (m *stateModel) pageDown() {
// 	m.screenStart += m.screenHeight
// 	m.cursor = m.screenStart
// 	m.clampScreenStart()
// 	m.clampCursor()
// }

func (m *stateModel) updateHelpView(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "?", "q":
			m.state = viewState
		}
	}
	return m, nil
}

// replace the `previousHeight`th lines under the cursor by the View of `by`
// while keeping the previous indentation
// this allows swapping the previous content of a Model by the new content
func (m *stateModel) replaceAtCursor(by render.Model, previousHeight int) {
	indentation := m.screen[m.cursor].Indentation
	nextLines := by.View()
	render.IndentBy(nextLines, indentation)
	m.screen = slices.Replace(m.screen, m.cursor, m.cursor+previousHeight, nextLines...)
}

func (m *stateModel) updateStateView(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			selected := m.Selected()
			if collapser, ok := selected.(render.Collapser); ok {
				previousHeight := len(selected.View())
				collapser.Expand()
				m.replaceAtCursor(selected, previousHeight)

				m.clampCursor()
				m.clampScreen()
			}
		case "backspace":
			selected := m.Selected()
			if collapser, ok := selected.(render.Collapser); ok {
				previousLines := selected.View()
				if m.screen[m.cursor].PointsToEnd {
					m.cursor -= len(previousLines) - 1
				}
				collapser.Collapse()
				m.replaceAtCursor(selected, len(previousLines))

				m.clampCursor()
				m.clampScreen()
			}
		case "r":
			selected := m.Selected()
			if sensitiveValue, ok := selected.(*json.SensitiveValue); ok {
				height := len(selected.View())
				sensitiveValue.Reveal()
				m.replaceAtCursor(selected, height)

				m.clampCursor()
				m.clampScreen()
			}
		}
	case tea.WindowSizeMsg:
		m.screenHeight = msg.Height
		m.screenWidth = msg.Width

		m.clampCursor()
		m.clampScreen()
	}
	return m, nil
}

func (m *stateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case showHelp:
		return m.updateHelpView(msg)
	case viewState:
		return m.updateStateView(msg)
	}
	panic("unknown state")
}
