package cmd

import (
	"github.com/Art-S-D/tfx/internal/node"
	"github.com/Art-S-D/tfx/internal/style"
	tea "github.com/charmbracelet/bubbletea"
)

type modelState int

const (
	viewState = modelState(iota)
	showHelp
)

type TfxModel struct {
	screenWidth, screenHeight int
	cursor                    *node.Node
	screenStart               *node.Node

	root  *node.Node
	state modelState
	theme *style.Theme
}

func (m *TfxModel) screenEnd() *node.Node {
	end := m.screenStart
	for i := 0; i < m.screenHeight-2 && end.Next() != nil; i++ {
		end = end.Next()
	}
	return end
}

func (m *TfxModel) Init() tea.Cmd {
	return tea.SetWindowTitle("tfx")
}

func NewModel(root *node.Node, theme *style.Theme) *TfxModel {
	out := &TfxModel{
		root:        root,
		theme:       theme,
		screenStart: root.Next(),
		cursor:      root.Next(),
	}
	return out
}
