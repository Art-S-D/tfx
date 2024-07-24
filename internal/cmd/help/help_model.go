package help

import (
	"fmt"

	tfxcontext "github.com/Art-S-D/tfx/internal/cmd/tfxcontext"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// missing keybindings from fx
// 1-9                 collapse to nth level
// z                   toggle string wrap
// y                   yank/copy

var helpText string = `
    Key Bindings
        q, ctrl+c, exc      exit program
        pgdown, space, f    page down
        pgup, b             page up
        u, ctrl+u           half page up
        d, ctrl+d    	    half page down
        g, home             go to top
        G, end              go to bottom
        down, j             down
        up, k               up
        ?                   show help
        right, l, enter     expand
        left, h, backspace  collapse
        L, shift+right      expand recursively
        H, shift+left       collapse recursively
        e                   expand all
        E                   collapse all
        J, shift+down       next sibling
        K, shift+up         previous sibling
        r                   reveal value
        /                   search regexp
        n                   next search result
        N                   prev search result
        p                   preview
        P                   print
        .                   dig
`

type HelpModel struct {
	Ctx         *tfxcontext.TfxContext
	ParentModel tea.Model
}

func (m *HelpModel) Init() tea.Cmd {
	return nil
}

func (m *HelpModel) View() string {
	help := lipgloss.PlaceVertical(m.Ctx.ScreenHeight-1, lipgloss.Top, helpText)

	lastLine := ": press q or ? to close help"
	lastLine = lipgloss.PlaceHorizontal(m.Ctx.ScreenWidth, lipgloss.Left, lastLine)
	lastLine = m.Ctx.Theme.Selection(lastLine).String()

	return fmt.Sprintf("%s\n%s", help, lastLine)
}

func (m *HelpModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.Ctx.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "?", "q", "esc":
			return m.ParentModel, nil
		}
	}
	return m, nil
}
