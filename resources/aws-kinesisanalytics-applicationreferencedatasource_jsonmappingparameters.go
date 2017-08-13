package resources

// AWS::KinesisAnalytics::ApplicationReferenceDataSource.JSONMappingParameters AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-jsonmappingparameters.html
type AWSKinesisAnalyticsApplicationReferenceDataSourceJSONMappingParameters struct {
    
    // RecordRowPath AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-jsonmappingparameters.html#cfn-kinesisanalytics-applicationreferencedatasource-jsonmappingparameters-recordrowpath
    RecordRowPath string `json:"RecordRowPath"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationReferenceDataSourceJSONMappingParameters) AWSCloudFormationType() string {
    return "AWS::KinesisAnalytics::ApplicationReferenceDataSource.JSONMappingParameters"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationReferenceDataSourceJSONMappingParameters) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}