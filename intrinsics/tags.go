package intrinsics

import (
	yaml "gopkg.in/yaml.v3"
	"strings"
)

var allTags = []string{
	"Ref", "GetAtt", "Base64", "FindInMap", "GetAZs",
	"ImportValue", "Join", "Select", "Split", "Sub",
	"Equals", "Cidr", "And", "If", "Not", "Or",
}

var tagResolvers = make(map[string]func(*yaml.Node) (*yaml.Node, error))

type fragment struct {
	content *yaml.Node
}

func (f *fragment) UnmarshalYAML(value *yaml.Node) error {
	var err error
	f.content, err = resolveTags(value)
	return err
}

type customTagProcessor struct {
	target interface{}
}

func (i *customTagProcessor) UnmarshalYAML(value *yaml.Node) error {
	resolved, err := resolveTags(value)
	if err != nil {
		return err
	}

	return resolved.Decode(i.target)
}

func addTagResolver(tag string, resolver func(*yaml.Node) (*yaml.Node, error)) {
	tagResolvers[tag] = resolver
}

func resolveTags(node *yaml.Node) (*yaml.Node, error) {
	for tag, fn := range tagResolvers {
		if node.Tag == tag {
			return fn(node)
		}
	}

	if node.Kind == yaml.SequenceNode || node.Kind == yaml.MappingNode {
		var err error
		for i := range node.Content {
			node.Content[i], err = resolveTags(node.Content[i])
			if err != nil {
				return nil, err
			}
		}
	}

	return node, nil
}

func registerTagMarshallers() {
	for _, tag := range allTags {
		addTagResolver("!"+tag, func(tag string) func(node *yaml.Node) (*yaml.Node, error) {
			return func(node *yaml.Node) (*yaml.Node, error) {
				prefix := "Fn::"
				if tag == "Ref" || tag == "Condition" || strings.HasPrefix(tag, prefix) {
					prefix = ""
				}

				tag = prefix + tag

				output := &yaml.Node{
					Kind:  yaml.MappingNode,
					Tag:   "!!map",
					Value: "",
					Content: []*yaml.Node{
						{Kind: yaml.ScalarNode, Tag: "!!str", Value: tag},
					},
				}

				if len(node.Content) == 0 {
					output.Content = append(output.Content, &yaml.Node{
						Kind:  yaml.ScalarNode,
						Tag:   "!!str",
						Value: node.Value,
					})
				} else {
					output.Content = append(output.Content, &yaml.Node{
						Kind:    yaml.SequenceNode,
						Tag:     "!!seq",
						Content: node.Content,
					})
				}

				return output, nil
			}
		}(tag))
	}
}
