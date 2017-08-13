package resources

// AWS::KinesisAnalytics::ApplicationOutput AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationoutput.html
type AWSKinesisAnalyticsApplicationOutput struct {
    
    // ApplicationName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationoutput.html#cfn-kinesisanalytics-applicationoutput-applicationname
    ApplicationName string `json:"ApplicationName"`
    
    // Output AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationoutput.html#cfn-kinesisanalytics-applicationoutput-output
    Output AWSKinesisAnalyticsApplicationOutputOutput `json:"Output"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationOutput) AWSCloudFormationType() string {
    return "AWS::KinesisAnalytics::ApplicationOutput"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationOutput) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}