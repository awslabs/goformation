package resources

// AWS::Logs::MetricFilter AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html
type AWSLogsMetricFilter struct {

	// FilterPattern AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-cwl-metricfilter-filterpattern

	FilterPattern string `json:"FilterPattern"`

	// LogGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-cwl-metricfilter-loggroupname

	LogGroupName string `json:"LogGroupName"`

	// MetricTransformations AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-cwl-metricfilter-metrictransformations

	MetricTransformations []AWSLogsMetricFilter_MetricTransformation `json:"MetricTransformations"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLogsMetricFilter) AWSCloudFormationType() string {
	return "AWS::Logs::MetricFilter"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLogsMetricFilter) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
