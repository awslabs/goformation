package cloudformation

// AWSEC2EC2Fleet_OnDemandOptionsRequest AWS CloudFormation Resource (AWS::EC2::EC2Fleet.OnDemandOptionsRequest)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ondemandoptionsrequest.html
type AWSEC2EC2Fleet_OnDemandOptionsRequest struct {

	// AllocationStrategy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ondemandoptionsrequest.html#cfn-ec2-ec2fleet-ondemandoptionsrequest-allocationstrategy
	AllocationStrategy string `json:"AllocationStrategy,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2EC2Fleet_OnDemandOptionsRequest) AWSCloudFormationType() string {
	return "AWS::EC2::EC2Fleet.OnDemandOptionsRequest"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSEC2EC2Fleet_OnDemandOptionsRequest) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
