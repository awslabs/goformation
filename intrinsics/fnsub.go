package intrinsics

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
)

// ResolveFnSub resolves the 'Fn::Sub' AWS CloudFormation intrinsic function.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-sub.html
func FnSub(name string, input interface{}, template interface{}) interface{} {

	// Input can either be a string for this type of Fn::Sub call:
	// { "Fn::Sub": "some-string-with-a-${variable}" }

	// or it will be an array of length two for named replacements
	// { "Fn::Sub": [ "some ${replaced}", { "replaced": "value" } ] }

	switch val := input.(type) {

	case []interface{}:
		// Replace each of the variables in element 0 with the items in element 1
		if src, ok := val[0].(string); ok {
			// The seconds element is a map of variables to replace
			if replacements, ok := val[1].(map[string]interface{}); ok {
				// Loop through the replacements
				for key, replacement := range replacements {
					// Check the replacement is a string
					if value, ok := replacement.(string); ok {
						src = strings.Replace(src, "${"+key+"}", value, -1)
					}
				}
			}
			// Replace rest with parameters
			regex := regexp.MustCompile(`\$\{.*\}`)
			variables := regex.FindAllStringSubmatch(src, -1)
			for _, variable := range variables {
				src = strings.Replace(src, variable[0], getParamValue(variable[0][2:len(variable[0])-1], template), -1)
			}
		}

	case string:
		// Look up references for each of the variables
		regex := regexp.MustCompile(`\$\{([\.0-9A-Za-z]+)\}`)
		variables := regex.FindAllStringSubmatch(val, -1)
		for _, variable := range variables {

			var resolved interface{}
			if strings.Contains(variable[1], ".") {
				// If the variable name has a . in it, use Fn::GetAtt to resolve it
				resolved = FnGetAtt("Fn::GetAtt", strings.Split(variable[1], "."), template)
			} else {
				// The variable name doesn't have a . in it, so use Ref
				resolved = Ref("Ref", variable[1], template)
			}

			if resolved != nil {
				if replacement, ok := resolved.(string); ok {
					val = strings.Replace(val, variable[0], replacement, -1)
				}
			} else {
				// The reference couldn't be resolved, so just strip the variable
				val = strings.Replace(val, variable[0], "", -1)
			}

		}

		// Replace rest with parameters
		regex = regexp.MustCompile(`\$\{.*\}`)
		variables = regex.FindAllStringSubmatch(val, -1)
		for _, variable := range variables {
			val = strings.Replace(val, variable[0], getParamValue(variable[0][2:len(variable[0])-1], template), -1)
		}
		return val
	}

	return nil
}

func getParamValue(name string, template interface{}) string {
	switch name {

	case "AWS::AccountId":
		return "123456789012"
	case "AWS::NotificationARNs": //
		return "arn:aws:sns:us-east-1:123456789012:MyTopic"
	case "AWS::NoValue":
		return ""
	case "AWS::Region":
		return "us-east-1"
	case "AWS::StackId":
		return "arn:aws:cloudformation:us-east-1:123456789012:stack/MyStack/1c2fa620-982a-11e3-aff7-50e2416294e0"
	case "AWS::StackName":
		return "goformation-stack"

	default:

		// This isn't a pseudo 'Ref' paramater, so we need to look inside the CloudFormation template
		// to see if we can resolve the reference. This implementation just looks at the Parameters section
		// to see if there is a parameter matching the name, and if so, return the default value.

		// Check the template is a map
		if template, ok := template.(map[string]interface{}); ok {
			// Check there is a parameters section
			if uparameters, ok := template["Parameters"]; ok {
				// Check the parameters section is a map
				if parameters, ok := uparameters.(map[string]interface{}); ok {
					// Check there is a parameter with the same name as the Ref
					if uparameter, ok := parameters[name]; ok {
						// Check the parameter is a map
						if parameter, ok := uparameter.(map[string]interface{}); ok {
							// Check the parameter has a default
							if def, ok := parameter["Default"]; ok {
								if src, ok := def.(string); ok {
									return src
								} else {
									return fmt.Sprintf("%v", def)
								}
							}
						}
					}
				}
			}
		}
	}

	return ""
}

// NewSub substitutes variables in an input string with values that you specify. In your templates, you can use this function to construct commands or outputs that include values that aren't available until you create or update a stack.
func Sub(value string) string {
	i := `{ "Fn::Sub" : "` + value + `" }`
	return base64.StdEncoding.EncodeToString([]byte(i))
}
