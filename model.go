package main

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
	"github.com/Art-S-D/tfx/internal/tfstate"
	tea "github.com/charmbracelet/bubbletea"
)

type modelState int

const (
	viewState = modelState(iota)
	showHelp
)

type stateModel struct {
	screenWidth, screenHeight int
	cursor                    int
	screenStart               int // should always be between [0, rootModule.Height() - screenHeight)

	rootModule *tfstate.RootModuleModel
	screen     []render.Line

	state modelState

	theme *style.Theme
}

func (m *stateModel) refreshScreen() {
	m.screen = m.rootModule.View()
}
func (m *stateModel) Init() tea.Cmd {
	return tea.SetWindowTitle("tfx")
}
