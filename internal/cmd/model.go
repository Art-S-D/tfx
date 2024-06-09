package cmd

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

type TfxModel struct {
	screenWidth, screenHeight int
	cursor                    int
	screenStart               int // should always be between [0, rootModule.Height() - screenHeight)

	rootModule *tfstate.RootModuleModel
	screen     []render.Line

	state modelState

	theme *style.Theme
}

func (m *TfxModel) refreshScreen() {
	m.screen = m.rootModule.View()
}
func (m *TfxModel) Init() tea.Cmd {
	return tea.SetWindowTitle("tfx")
}

func NewModel(rootModule *tfstate.RootModuleModel, theme *style.Theme) *TfxModel {
	out := &TfxModel{
		rootModule: rootModule,
		theme:      theme,
	}
	out.refreshScreen()
	return out
}
