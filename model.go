package main

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
	"github.com/Art-S-D/tfx/internal/tfstate"
	tea "github.com/charmbracelet/bubbletea"
)

type stateModel struct {
	screenWidth, screenHeight int
	cursor                    int
	screenStart               int // should always be between [0, rootModule.Height() - screenHeight)

	rootModule *tfstate.RootModuleModel
	screen     []render.Line

	theme *style.Theme
}

func (m *stateModel) RefreshScreen() {
	m.screen = m.rootModule.View()
}
func (m *stateModel) Init() tea.Cmd {
	m.RefreshScreen()
	return tea.SetWindowTitle("tfx")
}
