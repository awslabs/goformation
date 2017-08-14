package resources

// AWS::EMR::InstanceGroupConfig.ScalingTrigger AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingtrigger.html
type AWSEMRInstanceGroupConfigScalingTrigger struct {

	// CloudWatchAlarmDefinition AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingtrigger.html#cfn-elasticmapreduce-instancegroupconfig-scalingtrigger-cloudwatchalarmdefinition
	CloudWatchAlarmDefinition AWSEMRInstanceGroupConfigScalingTriggerCloudWatchAlarmDefinition `json:"CloudWatchAlarmDefinition"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfigScalingTrigger) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.ScalingTrigger"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfigScalingTrigger) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
