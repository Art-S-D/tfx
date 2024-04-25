package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Art-S-D/tfview/internal/style"
	tea "github.com/charmbracelet/bubbletea"
	tfjson "github.com/hashicorp/terraform-json"
)

type stateModel struct {
	screenWidth, screenHeight int
	cursor                    int
	offset                    int // should always be between [0, rootModule.Height() - screenHeight)
	rootModule                StateModuleModel
	rootModuleHeight          int
}

func (m *stateModel) Init() tea.Cmd {
	m.rootModuleHeight = m.rootModule.Height()
	return nil
}
func (m *stateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			m.cursorUp()
			return m, nil
		case "down":
			m.cursorDown()
			return m, nil
		case "enter":
			selected := m.rootModule.Selected(m.cursor)
			selected.Toggle()
			m.rootModuleHeight = m.rootModule.Height()
			m.clampOffset()
		}
	case tea.WindowSizeMsg:
		m.screenHeight = msg.Height
		m.screenWidth = msg.Width

		m.clampOffset()
	}
	return m, nil
}

func (m *stateModel) View() string {
	view, _ := m.rootModule.ViewCursor(m.cursor)
	lines := strings.Split(view, "\n")
	linesInView := lines[m.offset : m.offset+m.screenHeight]
	if len(linesInView) > 1 {
		selection := m.rootModule.Selected(m.cursor)
		selectionName := selection.Address()
		linesInView[len(linesInView)-1] = style.Selection.Copy().
			PaddingRight(m.screenWidth - len(selectionName)).
			Render(selectionName)
	}
	return strings.Join(linesInView, "\n")
}

func (m *stateModel) clampOffset() {
	if m.offset < 0 {
		m.offset = 0
	}
	if m.offset >= m.rootModuleHeight-m.screenHeight {
		m.offset = m.rootModuleHeight - m.screenHeight - 1
	}
}

func (m *stateModel) cursorUp() {
	m.cursor -= 1
	if m.cursor < 0 {
		m.cursor = 0
	}
	if m.cursor <= m.offset+3 {
		m.offset -= 1
	}
	m.clampOffset()
}
func (m *stateModel) cursorDown() {
	m.cursor += 1
	// should be -1 but we have -2 instead to account for the preview
	screenBottom := m.screenHeight + m.offset - 2
	if m.cursor >= screenBottom {
		m.cursor = screenBottom - 1
	}
	if m.cursor >= screenBottom-3 {
		m.offset += 1
	}
	m.clampOffset()
}

func main() {
	stateFile, err := os.ReadFile("state.json")
	if err != nil {
		panic(err.Error())
	}
	var plan tfjson.State
	json.Unmarshal(stateFile, &plan)

	// resourceDef := `{
	// 	"address": "aws_ssm_parameter.dns_forwarding_vpc_id",
	// 	"mode": "managed",
	// 	"type": "aws_ssm_parameter",
	// 	"name": "dns_forwarding_vpc_id",
	// 	"provider_name": "registry.terraform.io/hashicorp/aws",
	// 	"schema_version": 0,
	// 	"values": {
	// 		"allowed_pattern": "",
	// 		"arn": "arn:aws:ssm:eu-west-3:899011411636:parameter/lz/tfvars/eu-xf/dns_forwarding_vpc_id_eu-west-1",
	// 		"data_type": "text",
	// 		"description": "",
	// 		"id": "/lz/tfvars/eu-xf/dns_forwarding_vpc_id_eu-west-1",
	// 		"insecure_value": null,
	// 		"key_id": "",
	// 		"name": "/lz/tfvars/eu-xf/dns_forwarding_vpc_id_eu-west-1",
	// 		"overwrite": null,
	// 		"tags": {},
	// 		"tags_all": {
	// 			"stla_lz": "true",
	// 			"stla_region": "eu-xf",
	// 			"stla_team": "ccoe_nw"
	// 		},
	// 		"test": [12, "aaa", {"othertest": []}],
	// 		"tier": "Standard",
	// 		"type": "String",
	// 		"value": "vpc-0c1205190b6218989",
	// 		"version": 1
	// 	},
	// 	"sensitive_values": {
	// 		"tags": {},
	// 		"tags_all": {},
	// 		"value": true
	// 	},
	// 	"depends_on": [
	// 		"data.aws_region.current",
	// 		"module.svc_prod_vpc.aws_vpc.svc"
	// 	]
	// }`
	// var resource tfjson.StateResource
	// err := json.Unmarshal([]byte(resourceDef), &resource)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// p := tea.NewProgram(StateResourceModel{resource, true})
	// if _, err := p.Run(); err != nil {
	// 	fmt.Printf("Alas, there's been an error: %v", err)
	// 	os.Exit(1)
	// }

	terraformState := stateModel{rootModule: StateModuleModelFromJson(*plan.Values.RootModule)}
	terraformState.rootModule.expanded = true
	p := tea.NewProgram(&terraformState)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	// fmt.Println(StateModuleModelFromJson(*plan.Values.RootModule))
}
