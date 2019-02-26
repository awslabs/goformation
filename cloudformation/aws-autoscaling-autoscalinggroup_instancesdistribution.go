package cloudformation

// AWSAutoScalingAutoScalingGroup_InstancesDistribution AWS CloudFormation Resource (AWS::AutoScaling::AutoScalingGroup.InstancesDistribution)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-mixedinstancespolicy-instancesdistribution.html
type AWSAutoScalingAutoScalingGroup_InstancesDistribution struct {

	// OnDemandAllocationStrategy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-mixedinstancespolicy-instancesdistribution.html#cfn-autoscaling-autoscalinggroup-instancesdistribution-ondemandallocationstrategy
	OnDemandAllocationStrategy string `json:"OnDemandAllocationStrategy,omitempty"`

	// OnDemandBaseCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-mixedinstancespolicy-instancesdistribution.html#cfn-autoscaling-autoscalinggroup-instancesdistribution-ondemandbasecapacity
	OnDemandBaseCapacity int `json:"OnDemandBaseCapacity,omitempty"`

	// OnDemandPercentageAboveBaseCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-mixedinstancespolicy-instancesdistribution.html#cfn-autoscaling-autoscalinggroup-instancesdistribution-ondemandpercentageabovebasecapacity
	OnDemandPercentageAboveBaseCapacity int `json:"OnDemandPercentageAboveBaseCapacity,omitempty"`

	// SpotAllocationStrategy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-mixedinstancespolicy-instancesdistribution.html#cfn-autoscaling-autoscalinggroup-instancesdistribution-spotallocationstrategy
	SpotAllocationStrategy string `json:"SpotAllocationStrategy,omitempty"`

	// SpotInstancePools AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-mixedinstancespolicy-instancesdistribution.html#cfn-autoscaling-autoscalinggroup-instancesdistribution-spotinstancepools
	SpotInstancePools int `json:"SpotInstancePools,omitempty"`

	// SpotMaxPrice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-mixedinstancespolicy-instancesdistribution.html#cfn-autoscaling-autoscalinggroup-instancesdistribution-spotmaxprice
	SpotMaxPrice string `json:"SpotMaxPrice,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingAutoScalingGroup_InstancesDistribution) AWSCloudFormationType() string {
	return "AWS::AutoScaling::AutoScalingGroup.InstancesDistribution"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAutoScalingAutoScalingGroup_InstancesDistribution) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
