package tfxcontext

import (
	"github.com/Art-S-D/tfx/internal/node"
	"github.com/Art-S-D/tfx/internal/style"
	tea "github.com/charmbracelet/bubbletea"
)

type TfxContext struct {
	ScreenWidth, ScreenHeight int
	Theme                     *style.Theme
	PrintOnExit               *node.Node
}

func (ctx *TfxContext) Update(msg tea.Msg) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		ctx.ScreenHeight = msg.Height
		ctx.ScreenWidth = msg.Width
	}
}
