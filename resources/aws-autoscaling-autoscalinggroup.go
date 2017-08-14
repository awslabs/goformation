package resources

// AWS::AutoScaling::AutoScalingGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html
type AWSAutoScalingAutoScalingGroup struct {

	// AvailabilityZones AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-availabilityzones
	AvailabilityZones []AWSAutoScalingAutoScalingGroupstring `json:"AvailabilityZones"`

	// Cooldown AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-cooldown
	Cooldown string `json:"Cooldown"`

	// DesiredCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-desiredcapacity
	DesiredCapacity string `json:"DesiredCapacity"`

	// HealthCheckGracePeriod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-healthcheckgraceperiod
	HealthCheckGracePeriod int64 `json:"HealthCheckGracePeriod"`

	// HealthCheckType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-healthchecktype
	HealthCheckType string `json:"HealthCheckType"`

	// InstanceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-instanceid
	InstanceId string `json:"InstanceId"`

	// LaunchConfigurationName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-launchconfigurationname
	LaunchConfigurationName string `json:"LaunchConfigurationName"`

	// LoadBalancerNames AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-loadbalancernames
	LoadBalancerNames []AWSAutoScalingAutoScalingGroupstring `json:"LoadBalancerNames"`

	// MaxSize AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-maxsize
	MaxSize string `json:"MaxSize"`

	// MetricsCollection AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-metricscollection
	MetricsCollection []AWSAutoScalingAutoScalingGroupMetricsCollection `json:"MetricsCollection"`

	// MinSize AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-minsize
	MinSize string `json:"MinSize"`

	// NotificationConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-notificationconfigurations
	NotificationConfigurations []AWSAutoScalingAutoScalingGroupNotificationConfiguration `json:"NotificationConfigurations"`

	// PlacementGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-placementgroup
	PlacementGroup string `json:"PlacementGroup"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-tags
	Tags []AWSAutoScalingAutoScalingGroupTagProperty `json:"Tags"`

	// TargetGroupARNs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-targetgrouparns
	TargetGroupARNs []AWSAutoScalingAutoScalingGroupstring `json:"TargetGroupARNs"`

	// TerminationPolicies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-termpolicy
	TerminationPolicies []AWSAutoScalingAutoScalingGroupstring `json:"TerminationPolicies"`

	// VPCZoneIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-vpczoneidentifier
	VPCZoneIdentifier []AWSAutoScalingAutoScalingGroupstring `json:"VPCZoneIdentifier"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingAutoScalingGroup) AWSCloudFormationType() string {
	return "AWS::AutoScaling::AutoScalingGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSAutoScalingAutoScalingGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
