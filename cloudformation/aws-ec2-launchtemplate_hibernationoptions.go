package cloudformation

// AWSEC2LaunchTemplate_HibernationOptions AWS CloudFormation Resource (AWS::EC2::LaunchTemplate.HibernationOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-hibernationoptions.html
type AWSEC2LaunchTemplate_HibernationOptions struct {

	// Configured AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-hibernationoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-hibernationoptions-configured
	Configured bool `json:"Configured,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2LaunchTemplate_HibernationOptions) AWSCloudFormationType() string {
	return "AWS::EC2::LaunchTemplate.HibernationOptions"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSEC2LaunchTemplate_HibernationOptions) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
