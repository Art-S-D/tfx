package main

import (
	"github.com/Art-S-D/tfview/internal/tfstate"
	tea "github.com/charmbracelet/bubbletea"
)

type stateModel struct {
	screenWidth, screenHeight int
	cursor                    int
	offset                    int // should always be between [0, rootModule.Height() - screenHeight)
	rootModule                tfstate.StateModuleModel
	rootModuleHeight          int

	// make the screen move with the cursor if the cursor is at a distance lower
	// than `screenDrag` to the top or the bottom of the screen
	screenDrag int
}

func (m *stateModel) clampOffset() {
	if m.offset < 0 {
		m.offset = 0
	}
	if m.offset >= m.rootModuleHeight-m.screenHeight {
		m.offset = m.rootModuleHeight - m.screenHeight - 1
	}
}

func (m *stateModel) screenBottom() int {
	// should be -1 but we have -2 instead to account for the preview
	return m.screenHeight + m.offset - 2
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

func (m *stateModel) Init() tea.Cmd {
	m.rootModuleHeight = m.rootModule.ViewHeight()
	return nil
}
