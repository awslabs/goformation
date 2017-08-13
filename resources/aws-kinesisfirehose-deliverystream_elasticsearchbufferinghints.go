package resources

// AWS::KinesisFirehose::DeliveryStream.ElasticsearchBufferingHints AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html
type AWSKinesisFirehoseDeliveryStreamElasticsearchBufferingHints struct {
    
    // IntervalInSeconds AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-bufferinghints-intervalinseconds
    IntervalInSeconds int64 `json:"IntervalInSeconds"`
    
    // SizeInMBs AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-bufferinghints-sizeinmbs
    SizeInMBs int64 `json:"SizeInMBs"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStreamElasticsearchBufferingHints) AWSCloudFormationType() string {
    return "AWS::KinesisFirehose::DeliveryStream.ElasticsearchBufferingHints"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStreamElasticsearchBufferingHints) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}