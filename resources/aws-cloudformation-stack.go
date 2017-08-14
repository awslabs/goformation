package resources

// AWS::CloudFormation::Stack AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html
type AWSCloudFormationStack struct {

	// NotificationARNs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html#cfn-cloudformation-stack-notificationarns
	NotificationARNs []string `json:"NotificationARNs"`

	// Parameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html#cfn-cloudformation-stack-parameters
	Parameters map[string]string `json:"Parameters"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html#cfn-cloudformation-stack-tags
	Tags []AWSCloudFormationStackTag `json:"Tags"`

	// TemplateURL AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html#cfn-cloudformation-stack-templateurl
	TemplateURL string `json:"TemplateURL"`

	// TimeoutInMinutes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html#cfn-cloudformation-stack-timeoutinminutes
	TimeoutInMinutes int64 `json:"TimeoutInMinutes"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFormationStack) AWSCloudFormationType() string {
	return "AWS::CloudFormation::Stack"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFormationStack) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
