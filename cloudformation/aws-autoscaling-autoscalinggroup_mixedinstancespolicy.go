package cloudformation

// AWSAutoScalingAutoScalingGroup_MixedInstancesPolicy AWS CloudFormation Resource (AWS::AutoScaling::AutoScalingGroup.MixedInstancesPolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-group-mixedinstancespolicy.html
type AWSAutoScalingAutoScalingGroup_MixedInstancesPolicy struct {

	// InstancesDistribution AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-group-mixedinstancespolicy.html#cfn-as-mixedinstancespolicy-instancesdistribution
	InstancesDistribution *AWSAutoScalingAutoScalingGroup_InstancesDistribution `json:"InstancesDistribution,omitempty"`

	// LaunchTemplate AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-group-mixedinstancespolicy.html#cfn-as-mixedinstancespolicy-launchtemplate
	LaunchTemplate *AWSAutoScalingAutoScalingGroup_LaunchTemplate `json:"LaunchTemplate,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingAutoScalingGroup_MixedInstancesPolicy) AWSCloudFormationType() string {
	return "AWS::AutoScaling::AutoScalingGroup.MixedInstancesPolicy"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAutoScalingAutoScalingGroup_MixedInstancesPolicy) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
