package resources

// AWS::Logs::MetricFilter.MetricTransformation AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html
type AWSLogsMetricFilterMetricTransformation struct {
    
    // MetricName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html#cfn-cwl-metricfilter-metrictransformation-metricname
    MetricName string `json:"MetricName"`
    
    // MetricNamespace AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html#cfn-cwl-metricfilter-metrictransformation-metricnamespace
    MetricNamespace string `json:"MetricNamespace"`
    
    // MetricValue AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html#cfn-cwl-metricfilter-metrictransformation-metricvalue
    MetricValue string `json:"MetricValue"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLogsMetricFilterMetricTransformation) AWSCloudFormationType() string {
    return "AWS::Logs::MetricFilter.MetricTransformation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLogsMetricFilterMetricTransformation) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}