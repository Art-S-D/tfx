package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Art-S-D/tfx/internal/cmd"
	"github.com/Art-S-D/tfx/internal/cmd/preview"
	tfxcontext "github.com/Art-S-D/tfx/internal/cmd/tfxcontext"
	"github.com/Art-S-D/tfx/internal/style"
	"github.com/Art-S-D/tfx/internal/tfstate"
	tea "github.com/charmbracelet/bubbletea"
	tfjson "github.com/hashicorp/terraform-json"
)

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		os.Remove("debug.log")
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	args := cmd.ParseArgs()

	jsonState, err := io.ReadAll(args.Src)
	if err != nil {
		panic(err.Error())
	}
	var plan tfjson.State
	err = json.Unmarshal(jsonState, &plan)
	if err != nil {
		panic(fmt.Errorf("failed to read json state %w", err))
	}

	context := &tfxcontext.TfxContext{Theme: style.DefaultTheme}
	rootPreview := &preview.PreviewModel{
		Ctx: context,
	}

	rootModule := tfstate.RootModuleNode(plan.Values.RootModule)
	if len(rootModule.Children()) == 0 {
		fmt.Println("The state is empty")
		return
	}
	rootPreview.SetRootNode(rootModule)

	p := tea.NewProgram(
		rootPreview,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	if _, err := p.Run(); err != nil {
		panic(err.Error())
	}

	if context.PrintOnExit != nil {
		fmt.Println(context.PrintOnExit.String())
	}
}
