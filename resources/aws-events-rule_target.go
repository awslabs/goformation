package resources

// AWSEventsRule_Target AWS CloudFormation Resource (AWS::Events::Rule.Target)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html
type AWSEventsRule_Target struct {

	// Arn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-arn
	Arn string `json:"Arn"`

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-id
	Id string `json:"Id"`

	// Input AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-input
	Input string `json:"Input"`

	// InputPath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-inputpath
	InputPath string `json:"InputPath"`

	// RoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-rolearn
	RoleArn string `json:"RoleArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEventsRule_Target) AWSCloudFormationType() string {
	return "AWS::Events::Rule.Target"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEventsRule_Target) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
