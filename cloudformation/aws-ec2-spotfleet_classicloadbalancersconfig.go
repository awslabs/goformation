package cloudformation

// AWSEC2SpotFleet_ClassicLoadBalancersConfig AWS CloudFormation Resource (AWS::EC2::SpotFleet.ClassicLoadBalancersConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-classicloadbalancersconfig.html
type AWSEC2SpotFleet_ClassicLoadBalancersConfig struct {

	// ClassicLoadBalancers AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-classicloadbalancersconfig.html#cfn-ec2-spotfleet-classicloadbalancersconfig-classicloadbalancers
	ClassicLoadBalancers []AWSEC2SpotFleet_ClassicLoadBalancer `json:"ClassicLoadBalancers,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_ClassicLoadBalancersConfig) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.ClassicLoadBalancersConfig"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSEC2SpotFleet_ClassicLoadBalancersConfig) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
