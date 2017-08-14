package resources

// AWS::KinesisFirehose::DeliveryStream AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html
type AWSKinesisFirehoseDeliveryStream struct {

	// DeliveryStreamName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverstream-deliverystreamname
	DeliveryStreamName string `json:"DeliveryStreamName"`

	// ElasticsearchDestinationConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverstream-elasticsearchdestinationconfiguration
	ElasticsearchDestinationConfiguration AWSKinesisFirehoseDeliveryStreamElasticsearchDestinationConfiguration `json:"ElasticsearchDestinationConfiguration"`

	// RedshiftDestinationConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-redshiftdestinationconfiguration
	RedshiftDestinationConfiguration AWSKinesisFirehoseDeliveryStreamRedshiftDestinationConfiguration `json:"RedshiftDestinationConfiguration"`

	// S3DestinationConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-s3destinationconfiguration
	S3DestinationConfiguration AWSKinesisFirehoseDeliveryStreamS3DestinationConfiguration `json:"S3DestinationConfiguration"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStream) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
