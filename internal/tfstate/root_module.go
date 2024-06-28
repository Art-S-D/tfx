package tfstate

import (
	"github.com/Art-S-D/tfx/internal/node"
	tfjson "github.com/hashicorp/terraform-json"
)

func RootModuleNode(json *tfjson.StateModule) *node.Node {
	root := &node.Node{}
	for _, resource := range json.Resources {
		childResource := StateResourceNode(resource)
		root.AppendChild(childResource)
	}
	for _, module := range json.ChildModules {
		childModule := StateModuleNode(module)
		root.AppendChild(childModule)
	}
	root.Expand()
	return root
}
