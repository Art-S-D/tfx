package main

import (
	"github.com/Art-S-D/tfx/internal/tfstate"
	tea "github.com/charmbracelet/bubbletea"
)

type stateModel struct {
	screenWidth, screenHeight int
	cursor                    int
	screenStart               int // should always be between [0, rootModule.Height() - screenHeight)

	rootModule       *tfstate.RootModuleModel
	rootModuleHeight int
}

func (m *stateModel) Init() tea.Cmd {
	m.rootModuleHeight = m.rootModule.ViewHeight()

	return tea.SetWindowTitle("tfx")
}
