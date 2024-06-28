package tfstate

import (
	"github.com/Art-S-D/tfx/internal/node"
	"github.com/Art-S-D/tfx/internal/style"
	tfjson "github.com/hashicorp/terraform-json"
)

func StateModuleNode(json *tfjson.StateModule) *node.Node {
	module := &node.Node{}
	module.SetAddress(json.Address)

	for _, resource := range json.Resources {
		childResource := StateResourceNode(resource)
		childResource.IncreaseDepth()
		module.AppendChild(childResource)
	}
	for _, mod := range json.ChildModules {
		childModule := StateModuleNode(mod)
		childModule.IncreaseDepth()
		module.AppendChild(childModule)
	}

	module.SetExpanded(expandedStateModule(json))
	module.SetCollapsed(collapsedStateModule(json))

	lastChild := node.String("}")
	lastChild.SetAddress(json.Address)
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
