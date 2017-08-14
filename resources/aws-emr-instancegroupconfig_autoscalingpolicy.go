package resources

// AWS::EMR::InstanceGroupConfig.AutoScalingPolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-autoscalingpolicy.html
type AWSEMRInstanceGroupConfigAutoScalingPolicy struct {

	// Constraints AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-autoscalingpolicy.html#cfn-elasticmapreduce-instancegroupconfig-autoscalingpolicy-constraints
	Constraints AWSEMRInstanceGroupConfigAutoScalingPolicyScalingConstraints `json:"Constraints"`

	// Rules AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-autoscalingpolicy.html#cfn-elasticmapreduce-instancegroupconfig-autoscalingpolicy-rules
	Rules []AWSEMRInstanceGroupConfigAutoScalingPolicyScalingRule `json:"Rules"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfigAutoScalingPolicy) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.AutoScalingPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfigAutoScalingPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
