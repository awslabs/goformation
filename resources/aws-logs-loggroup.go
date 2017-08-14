package resources

// AWS::Logs::LogGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html
type AWSLogsLogGroup struct {

	// LogGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-cwl-loggroup-loggroupname
	LogGroupName string `json:"LogGroupName"`

	// RetentionInDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-cwl-loggroup-retentionindays
	RetentionInDays int64 `json:"RetentionInDays"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLogsLogGroup) AWSCloudFormationType() string {
	return "AWS::Logs::LogGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLogsLogGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
