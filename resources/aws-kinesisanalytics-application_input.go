package resources

// AWS::KinesisAnalytics::Application.Input AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html
type AWSKinesisAnalyticsApplicationInput struct {
    
    // InputParallelism AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html#cfn-kinesisanalytics-application-input-inputparallelism
    InputParallelism AWSKinesisAnalyticsApplicationInputInputParallelism `json:"InputParallelism"`
    
    // InputSchema AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html#cfn-kinesisanalytics-application-input-inputschema
    InputSchema AWSKinesisAnalyticsApplicationInputInputSchema `json:"InputSchema"`
    
    // KinesisFirehoseInput AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html#cfn-kinesisanalytics-application-input-kinesisfirehoseinput
    KinesisFirehoseInput AWSKinesisAnalyticsApplicationInputKinesisFirehoseInput `json:"KinesisFirehoseInput"`
    
    // KinesisStreamsInput AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html#cfn-kinesisanalytics-application-input-kinesisstreamsinput
    KinesisStreamsInput AWSKinesisAnalyticsApplicationInputKinesisStreamsInput `json:"KinesisStreamsInput"`
    
    // NamePrefix AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html#cfn-kinesisanalytics-application-input-nameprefix
    NamePrefix string `json:"NamePrefix"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationInput) AWSCloudFormationType() string {
    return "AWS::KinesisAnalytics::Application.Input"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationInput) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}