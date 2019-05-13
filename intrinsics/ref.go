package intrinsics

// Ref resolves the 'Ref' AWS CloudFormation intrinsic function.
// Currently, this only resolves against CloudFormation Parameter default values
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html
func Ref(name string, input interface{}, template interface{}) interface{} {

	// Dang son, this has got more nest than a bald eagle
	// Check the input is a string
	if name, ok := input.(string); ok {

		switch name {

		case "AWS::AccountId":
			return "123456789012"
		case "AWS::NotificationARNs": //
			return []string{"arn:aws:sns:us-east-1:123456789012:MyTopic"}
		case "AWS::NoValue":
			return nil
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
									return def
								}
							}
						}
					}
				}
				if uresources, ok := template["Resources"]; ok {
					// Check the resources section is a map
					if resources, ok := uresources.(map[string]interface{}); ok {
						// Check there is a resource with the same name as the Ref
						if uresource, ok := resources[name]; ok {
							// Check the resource is a map
							if resource, ok := uresource.(map[string]interface{}); ok {
								// Check the resource type
								if _type, ok := resource["Type"]; ok {
									switch _type {
									case "AWS::Serverless::Function":
										if uprops, ok := resource["Properties"]; ok {
											if props, ok := uprops.(map[string]interface{}); ok {
												if funcName, ok := props["FunctionName"]; ok {
													switch v := funcName.(type) {
													case string:
														return v
													case map[string]interface{}:
														for key, value := range v {
															if key == "Fn::Sub" {
																return FnSub(key, value, v)
															}
														}
														return v
													}
												} else {
													return name
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

	}

	return nil

}
