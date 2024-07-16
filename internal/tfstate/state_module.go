package tfstate

import (
	"github.com/Art-S-D/tfx/internal/node"
	"github.com/Art-S-D/tfx/internal/style"
	tfjson "github.com/hashicorp/terraform-json"
)

func StateModuleNode(json *tfjson.StateModule) *node.Node {
	module := &node.Node{
		Address:  json.Address,
		Childer:  &node.DefaultChilder{},
		Expander: &node.DefaultExpander{},
		Renderer: &node.LineRenderer{
			Expanded:  expandedStateModule(json),
			Collapsed: collapsedStateModule(json),
		},
	}

	for _, resource := range json.Resources {
		childResource := StateResourceNode(resource)
		childResource.Indent()
		module.AppendChild(childResource)
	}
	for _, mod := range json.ChildModules {
		childModule := StateModuleNode(mod)
		childModule.Indent()
		module.AppendChild(childModule)
	}

	lastChild := node.StringNode("}")
	lastChild.Address = json.Address
	module.AppendChild(lastChild)

	return module
}

func expandedStateModule(module *tfjson.StateModule) style.Str {
	return style.Concat(
		style.Type("module").Selectable(),
		style.Space().Selectable(),
		style.Key(module.Address).Selectable(),
		style.Space().UnSelectable(),
		style.Default("{").UnSelectable(),
	)
}
func collapsedStateModule(module *tfjson.StateModule) style.Str {
	return style.Concat(
		expandedStateModule(module),
		style.Preview("...").UnSelectable(),
		style.Default("}").UnSelectable(),
	)
}
