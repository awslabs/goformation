package resources

// AWS::KinesisAnalytics::ApplicationReferenceDataSource.ReferenceDataSource AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html
type AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource struct {

	// ReferenceSchema AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-referencedatasource-referenceschema

	ReferenceSchema AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchema `json:"ReferenceSchema"`

	// S3ReferenceDataSource AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-referencedatasource-s3referencedatasource

	S3ReferenceDataSource AWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSource `json:"S3ReferenceDataSource"`

	// TableName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-referencedatasource-tablename

	TableName string `json:"TableName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationReferenceDataSource.ReferenceDataSource"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
