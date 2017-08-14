package resources

// AWS::KinesisAnalytics::Application.JSONMappingParameters AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-jsonmappingparameters.html
type AWSKinesisAnalyticsApplicationJSONMappingParameters struct {

	// RecordRowPath AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-jsonmappingparameters.html#cfn-kinesisanalytics-application-jsonmappingparameters-recordrowpath
	RecordRowPath string `json:"RecordRowPath"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationJSONMappingParameters) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.JSONMappingParameters"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationJSONMappingParameters) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
