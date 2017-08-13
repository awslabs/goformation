package resources

// AWS::KinesisAnalytics::ApplicationOutput.KinesisStreamsOutput AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-kinesisstreamsoutput.html
type AWSKinesisAnalyticsApplicationOutputKinesisStreamsOutput struct {
    
    // ResourceARN AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-kinesisstreamsoutput.html#cfn-kinesisanalytics-applicationoutput-kinesisstreamsoutput-resourcearn
    ResourceARN string `json:"ResourceARN"`
    
    // RoleARN AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-kinesisstreamsoutput.html#cfn-kinesisanalytics-applicationoutput-kinesisstreamsoutput-rolearn
    RoleARN string `json:"RoleARN"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationOutputKinesisStreamsOutput) AWSCloudFormationType() string {
    return "AWS::KinesisAnalytics::ApplicationOutput.KinesisStreamsOutput"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationOutputKinesisStreamsOutput) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}