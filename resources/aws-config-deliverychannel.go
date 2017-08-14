package resources

// AWS::Config::DeliveryChannel AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html
type AWSConfigDeliveryChannel struct {

	// ConfigSnapshotDeliveryProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-configsnapshotdeliveryproperties

	ConfigSnapshotDeliveryProperties AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties `json:"ConfigSnapshotDeliveryProperties"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-name

	Name string `json:"Name"`

	// S3BucketName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-s3bucketname

	S3BucketName string `json:"S3BucketName"`

	// S3KeyPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-s3keyprefix

	S3KeyPrefix string `json:"S3KeyPrefix"`

	// SnsTopicARN AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-snstopicarn

	SnsTopicARN string `json:"SnsTopicARN"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigDeliveryChannel) AWSCloudFormationType() string {
	return "AWS::Config::DeliveryChannel"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSConfigDeliveryChannel) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
