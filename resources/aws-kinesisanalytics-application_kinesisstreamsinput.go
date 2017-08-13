package resources

// AWS::KinesisAnalytics::Application.KinesisStreamsInput AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-kinesisstreamsinput.html
type AWSKinesisAnalyticsApplicationKinesisStreamsInput struct {
    
    // ResourceARN AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-kinesisstreamsinput.html#cfn-kinesisanalytics-application-kinesisstreamsinput-resourcearn
    ResourceARN string `json:"ResourceARN"`
    
    // RoleARN AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-kinesisstreamsinput.html#cfn-kinesisanalytics-application-kinesisstreamsinput-rolearn
    RoleARN string `json:"RoleARN"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationKinesisStreamsInput) AWSCloudFormationType() string {
    return "AWS::KinesisAnalytics::Application.KinesisStreamsInput"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationKinesisStreamsInput) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}