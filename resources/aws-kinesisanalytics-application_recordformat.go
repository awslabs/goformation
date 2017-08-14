package resources

// AWS::KinesisAnalytics::Application.RecordFormat AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordformat.html
type AWSKinesisAnalyticsApplicationRecordFormat struct {

	// MappingParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordformat.html#cfn-kinesisanalytics-application-recordformat-mappingparameters
	MappingParameters AWSKinesisAnalyticsApplicationRecordFormatMappingParameters `json:"MappingParameters"`

	// RecordFormatType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordformat.html#cfn-kinesisanalytics-application-recordformat-recordformattype
	RecordFormatType string `json:"RecordFormatType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationRecordFormat) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.RecordFormat"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationRecordFormat) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
