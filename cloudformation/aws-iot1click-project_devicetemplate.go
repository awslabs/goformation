package cloudformation

// AWSIoT1ClickProject_DeviceTemplate AWS CloudFormation Resource (AWS::IoT1Click::Project.DeviceTemplate)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot1click-project-devicetemplate.html
type AWSIoT1ClickProject_DeviceTemplate struct {

	// CallbackOverrides AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot1click-project-devicetemplate.html#cfn-iot1click-project-devicetemplate-callbackoverrides
	CallbackOverrides interface{} `json:"CallbackOverrides,omitempty"`

	// DeviceType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot1click-project-devicetemplate.html#cfn-iot1click-project-devicetemplate-devicetype
	DeviceType string `json:"DeviceType,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoT1ClickProject_DeviceTemplate) AWSCloudFormationType() string {
	return "AWS::IoT1Click::Project.DeviceTemplate"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSIoT1ClickProject_DeviceTemplate) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
