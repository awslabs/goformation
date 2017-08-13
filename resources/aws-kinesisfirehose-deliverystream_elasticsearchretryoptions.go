package resources

// AWS::KinesisFirehose::DeliveryStream.ElasticsearchRetryOptions AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-retryoptions.html
type AWSKinesisFirehoseDeliveryStreamElasticsearchRetryOptions struct {
    
    // DurationInSeconds AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-retryoptions.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-retryoptions-durationinseconds
    DurationInSeconds int64 `json:"DurationInSeconds"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStreamElasticsearchRetryOptions) AWSCloudFormationType() string {
    return "AWS::KinesisFirehose::DeliveryStream.ElasticsearchRetryOptions"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStreamElasticsearchRetryOptions) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}