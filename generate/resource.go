package main

// Resource represents an AWS CloudFormation resource such as AWS::EC2::Instance
type Resource struct {

	// Documentation is a link to the AWS CloudFormation User Guide for information about the resource.
	Documentation string `json:"Documentation"`

	// Properties are a list of property specifications for the resource. For details, see:
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-resource-specification-format.html#cfn-resource-specification-format-propertytypes
	Properties map[string]Property
}
