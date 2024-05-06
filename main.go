package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/Art-S-D/tfx/internal/tfstate"
	tea "github.com/charmbracelet/bubbletea"
	tfjson "github.com/hashicorp/terraform-json"
)

func main() {
	args := parseArgs()

	jsonState, err := io.ReadAll(args.src)
	if err != nil {
		panic(err.Error())
	}
	var plan tfjson.State
	err = json.Unmarshal(jsonState, &plan)
	if err != nil {
		panic(fmt.Errorf("failed to read json state %w", err))
	}

	terraformState := stateModel{rootModule: tfstate.RootModuleModelFromJson(*plan.Values.RootModule)}

	p := tea.NewProgram(&terraformState)
	if _, err := p.Run(); err != nil {
		panic(err.Error())
	}
}
