package cloudformation

// AWSEC2LaunchTemplate_LaunchTemplateElasticInferenceAccelerator AWS CloudFormation Resource (AWS::EC2::LaunchTemplate.LaunchTemplateElasticInferenceAccelerator)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplateelasticinferenceaccelerator.html
type AWSEC2LaunchTemplate_LaunchTemplateElasticInferenceAccelerator struct {

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplateelasticinferenceaccelerator.html#cfn-ec2-launchtemplate-launchtemplateelasticinferenceaccelerator-type
	Type string `json:"Type,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2LaunchTemplate_LaunchTemplateElasticInferenceAccelerator) AWSCloudFormationType() string {
	return "AWS::EC2::LaunchTemplate.LaunchTemplateElasticInferenceAccelerator"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSEC2LaunchTemplate_LaunchTemplateElasticInferenceAccelerator) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
