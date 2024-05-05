package main

import (
	"github.com/Art-S-D/tfview/internal/tfstate"
	tea "github.com/charmbracelet/bubbletea"
)

type stateModel struct {
	screenWidth, screenHeight int
	cursor                    int
	screenStart               int // should always be between [0, rootModule.Height() - screenHeight)
	rootModule                *tfstate.RootModuleModel
	rootModuleHeight          int

	// make the screen move with the cursor if the cursor is at a distance lower
	// than `screenDrag` to the top or the bottom of the screen
	screenDrag int
}

func (m *stateModel) screenBottom() int {
	// should be -1 but we have -2 instead to account for the preview
	return m.screenHeight + m.screenStart - 1
}

func (m *stateModel) Init() tea.Cmd {
	m.rootModuleHeight = m.rootModule.ViewHeight()

	return nil
}
