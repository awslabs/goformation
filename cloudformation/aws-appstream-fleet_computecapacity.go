package cloudformation

// AWSAppStreamFleet_ComputeCapacity AWS CloudFormation Resource (AWS::AppStream::Fleet.ComputeCapacity)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-computecapacity.html
type AWSAppStreamFleet_ComputeCapacity struct {

	// DesiredInstances AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-computecapacity.html#cfn-appstream-fleet-computecapacity-desiredinstances
	DesiredInstances int `json:"DesiredInstances,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppStreamFleet_ComputeCapacity) AWSCloudFormationType() string {
	return "AWS::AppStream::Fleet.ComputeCapacity"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAppStreamFleet_ComputeCapacity) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
