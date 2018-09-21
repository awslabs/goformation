package cloudformation

// AWSEC2SpotFleet_TargetGroup AWS CloudFormation Resource (AWS::EC2::SpotFleet.TargetGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-targetgroup.html
type AWSEC2SpotFleet_TargetGroup struct {

	// Arn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-targetgroup.html#cfn-ec2-spotfleet-targetgroup-arn
	Arn string `json:"Arn,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_TargetGroup) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.TargetGroup"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSEC2SpotFleet_TargetGroup) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
