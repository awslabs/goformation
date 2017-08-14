package resources

// AWS::Logs::LogStream AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html
type AWSLogsLogStream struct {

	// LogGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html#cfn-logs-logstream-loggroupname

	LogGroupName string `json:"LogGroupName"`

	// LogStreamName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html#cfn-logs-logstream-logstreamname

	LogStreamName string `json:"LogStreamName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLogsLogStream) AWSCloudFormationType() string {
	return "AWS::Logs::LogStream"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLogsLogStream) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
