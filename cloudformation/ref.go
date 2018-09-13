package cloudformation

import (
	"encoding/json"
)

// Ref creates a CloudFormation Reference to another resource in the template
func Ref(logicalName string) string {
	return `{ "Ref": "` + logicalName + `" }`
}

// GetAtt returns the value of an attribute from a resource in the template.
func GetAtt(logicalName string, attribute string) string {
	return `{ "Fn::GetAtt": [ "` + logicalName + `", "` + attribute + `" ] }`
}

// ImportValue returns the value of an output exported by another stack. You typically use this function to create cross-stack references. In the following example template snippets, Stack A exports VPC security group values and Stack B imports them.
func ImportValue(name string) string {
	return `{ "Fn::Import": "` + name + `" }`
}

// processReferences is a post processor that replaces all goformation references
// with proper CloudFormation references
func processReferences(input interface{}) (interface{}, error) {

	// Marshal to JSON and back to convert from a typed template object to simple primitives
	b, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	// Recurse through the object tree, replacing any Goformation references
	return replaceReferencesRecursive(m), nil

}

// replaceReferencesRecursive recurses through an object, and replaces any strings that
// contain '%%Ref:(.*)%%' with a proper AWS CloudFormation reference object
func replaceReferencesRecursive(input interface{}) interface{} {

	switch value := input.(type) {

	case map[string]interface{}:
		result := map[string]interface{}{}
		for k, v := range value {
			result[k] = replaceReferencesRecursive(v)
		}
		return result

	case []interface{}:
		result := []interface{}{}
		for _, v := range value {
			result = append(result, replaceReferencesRecursive(v))
		}
		return result

	case string:

		// Check if the string can be unmarshalled into an intrinsic object
		var intrinsic map[string]interface{}
		if err := json.Unmarshal([]byte(value), &intrinsic); err != nil {
			// The string value is not JSON, so it's not an intrinsic so just pass it back
			return value
		}

		// An intrinsic should be an object, with a single key containing a valid intrinsic name
		if len(intrinsic) != 1 {
			return value
		}

		supported := []string{
			"Ref",
			"Fn::Base64",
			"Fn::Cidr",
			"Fn::FindInMap",
			"Fn::GetAtt",
			"Fn::GetAZs",
			"Fn::ImportValue",
			"Fn::Join",
			"Fn::Select",
			"Fn::Split",
			"Fn::Sub",
			"Fn::Transform",
			"Fn::And",
			"Fn::Equals",
			"Fn::If",
			"Fn::Not",
			"Fn::Or",
		}

		for name, v := range intrinsic {
			for _, i := range supported {
				if name == i {
					return map[string]interface{}{
						name: replaceReferencesRecursive(v),
					}
				}
			}
		}

		return value

	default:
		return value
	}

}
