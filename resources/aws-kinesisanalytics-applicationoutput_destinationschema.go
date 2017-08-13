package resources

// AWS::KinesisAnalytics::ApplicationOutput.DestinationSchema AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-destinationschema.html
type AWSKinesisAnalyticsApplicationOutputDestinationSchema struct {
    
    // RecordFormatType AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-destinationschema.html#cfn-kinesisanalytics-applicationoutput-destinationschema-recordformattype
    RecordFormatType string `json:"RecordFormatType"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationOutputDestinationSchema) AWSCloudFormationType() string {
    return "AWS::KinesisAnalytics::ApplicationOutput.DestinationSchema"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationOutputDestinationSchema) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}