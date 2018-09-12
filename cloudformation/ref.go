package cloudformation

import (
	"encoding/json"
	"regexp"
)

// Ref creates a CloudFormation Reference to another resource in the template
func Ref(logicalName string) string {
	return "%%Ref:" + logicalName + "%%"
}

// GetAtt creates a CloudFormation reference to an attribute of another resource in the template
func GetAtt(logicalName string, attribute string) string {
	return "%%Fn::GetAtt:" + logicalName + "." + attribute + "%%"
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

		// Check if the string contains a goformation reference
		refexp := regexp.MustCompile("%%Ref:(.*)%%")
		refmatches := refexp.FindStringSubmatch(value)
		if len(refmatches) > 0 {
			return map[string]string{
				"Ref": refmatches[1],
			}
		}

		// Check if the string containers a goformation GetAtt reference
		getattexp := regexp.MustCompile(`%%Fn::GetAtt:(.*)\.(.*)%%`)
		getattmatches := getattexp.FindStringSubmatch(value)
		if len(getattmatches) > 0 {
			return map[string][]string{
				"Fn::GetAtt": []string{getattmatches[1], getattmatches[2]},
			}
		}
		return value

	default:
		return value

	}

}
