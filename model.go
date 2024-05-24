package main

import (
	"github.com/Art-S-D/tfx/internal/render"
	"github.com/Art-S-D/tfx/internal/style"
	tea "github.com/charmbracelet/bubbletea"
)

type stateModel struct {
	screenWidth, screenHeight int
	cursor                    int
	screenStart               int // should always be between [0, rootModule.Height() - screenHeight)

	rootNode *render.Node

	theme *style.Theme
}

func (m *stateModel) ViewParams() render.ViewParams {
	return render.ViewParams{
		CurrentLine:              0,
		Indentation:              0,
		SkipCursorForCurrentLine: false,

		Cursor:       m.cursor,
		ScreenStart:  m.screenStart,
		ScreenWidth:  m.screenWidth,
		ScreenHeight: m.screenHeight - 1,
	}
}

func (m *stateModel) Init() tea.Cmd {
	return tea.SetWindowTitle("tfx")
}
