package resources

// AWS::KinesisAnalytics::Application.InputParallelism AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputparallelism.html
type AWSKinesisAnalyticsApplicationInputParallelism struct {

	// Count AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputparallelism.html#cfn-kinesisanalytics-application-inputparallelism-count
	Count int64 `json:"Count"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationInputParallelism) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.InputParallelism"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationInputParallelism) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
