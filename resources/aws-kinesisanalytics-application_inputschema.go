package resources

// AWS::KinesisAnalytics::Application.InputSchema AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputschema.html
type AWSKinesisAnalyticsApplicationInputSchema struct {
    
    // RecordColumns AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputschema.html#cfn-kinesisanalytics-application-inputschema-recordcolumns
    RecordColumns []AWSKinesisAnalyticsApplicationInputSchemaRecordColumn `json:"RecordColumns"`
    
    // RecordEncoding AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputschema.html#cfn-kinesisanalytics-application-inputschema-recordencoding
    RecordEncoding string `json:"RecordEncoding"`
    
    // RecordFormat AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputschema.html#cfn-kinesisanalytics-application-inputschema-recordformat
    RecordFormat AWSKinesisAnalyticsApplicationInputSchemaRecordFormat `json:"RecordFormat"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationInputSchema) AWSCloudFormationType() string {
    return "AWS::KinesisAnalytics::Application.InputSchema"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationInputSchema) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}