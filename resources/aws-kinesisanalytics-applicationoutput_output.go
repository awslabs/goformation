package resources

// AWS::KinesisAnalytics::ApplicationOutput.Output AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html
type AWSKinesisAnalyticsApplicationOutputOutput struct {
    
    // DestinationSchema AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-destinationschema
    DestinationSchema AWSKinesisAnalyticsApplicationOutputOutputDestinationSchema `json:"DestinationSchema"`
    
    // KinesisFirehoseOutput AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-kinesisfirehoseoutput
    KinesisFirehoseOutput AWSKinesisAnalyticsApplicationOutputOutputKinesisFirehoseOutput `json:"KinesisFirehoseOutput"`
    
    // KinesisStreamsOutput AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-kinesisstreamsoutput
    KinesisStreamsOutput AWSKinesisAnalyticsApplicationOutputOutputKinesisStreamsOutput `json:"KinesisStreamsOutput"`
    
    // Name AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-name
    Name string `json:"Name"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationOutputOutput) AWSCloudFormationType() string {
    return "AWS::KinesisAnalytics::ApplicationOutput.Output"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationOutputOutput) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}