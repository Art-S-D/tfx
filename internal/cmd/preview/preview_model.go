package preview

import (
	tfxcontext "github.com/Art-S-D/tfx/internal/cmd/tfxcontext"
	"github.com/Art-S-D/tfx/internal/node"
	tea "github.com/charmbracelet/bubbletea"
)

type PreviewModel struct {
	Ctx *tfxcontext.TfxContext

	node        *node.Node
	cursor      *node.Node
	screenStart *node.Node

	ParentModel *PreviewModel
	PrintOnExit *node.Node
}

func (m *PreviewModel) SetNode(n *node.Node) {
	m.node = n
	m.cursor = n
	m.screenStart = n
}
func (m *PreviewModel) SetRootNode(n *node.Node) {
	m.SetNode(n)
	m.cursor = n.Next()
}

func (m *PreviewModel) Init() tea.Cmd {
	return nil
}
