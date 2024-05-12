package main

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/tfstate"
	tea "github.com/charmbracelet/bubbletea"
)

type stateModel struct {
	screenWidth, screenHeight int
	cursor                    int
	screenStart               int // should always be between [0, rootModule.Height() - screenHeight)

	screen []*render.ScreenLine

	rootModule *tfstate.RootModuleModel
}

func (m *stateModel) Init() tea.Cmd {
	return tea.SetWindowTitle("tfx")
}

func (m *stateModel) renderScreen() {
	m.screen = m.rootModule.Lines(0)
}

func (m *stateModel) totalHeight() int {
	return len(m.screen)
}

func (m *stateModel) selected() render.Model {
	return m.screen[m.cursor].PointsTo
}
