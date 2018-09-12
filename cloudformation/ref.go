package cloudformation

import (
	"encoding/json"
	"regexp"
)

// Ref creates a CloudFormation Reference to another resource in the template
func Ref(logicalName string) string {
	return "%%Ref:" + logicalName + "%%"
}

// updateReferences is a post processor that replaces all goformation references
// with proper CloudFormation references
func updateReferences(input interface{}) (interface{}, error) {

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
	return updateReferencesRecursive(m), nil

}

// updateReferencesRecursive
func updateReferencesRecursive(input interface{}) interface{} {

	switch value := input.(type) {

	case map[string]interface{}:
		result := map[string]interface{}{}
		for k, v := range value {
			result[k] = updateReferencesRecursive(v)
		}
		return result

	case []interface{}:
		result := []interface{}{}
		for _, v := range value {
			result = append(result, updateReferencesRecursive(v))
		}
		return result

	case string:
		// Check if the string contains a goformation reference
		re := regexp.MustCompile("%%Ref:(.*)%%")
		matches := re.FindStringSubmatch(value)
		if len(matches) > 0 {
			return map[string]string{
				"Ref": matches[1],
			}
		}
		return value

	default:
		return value

	}

}
