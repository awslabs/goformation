package resources

// AWSAutoScalingAutoScalingGroup_MetricsCollection AWS CloudFormation Resource (AWS::AutoScaling::AutoScalingGroup.MetricsCollection)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html
type AWSAutoScalingAutoScalingGroup_MetricsCollection struct {

	// Granularity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html#cfn-as-metricscollection-granularity
	Granularity string `json:"Granularity"`

	// Metrics AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html#cfn-as-metricscollection-metrics
	Metrics []string `json:"Metrics"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingAutoScalingGroup_MetricsCollection) AWSCloudFormationType() string {
	return "AWS::AutoScaling::AutoScalingGroup.MetricsCollection"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSAutoScalingAutoScalingGroup_MetricsCollection) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
