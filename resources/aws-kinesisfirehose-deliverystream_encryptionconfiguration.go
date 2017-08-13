package resources

// AWS::KinesisFirehose::DeliveryStream.EncryptionConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration.html
type AWSKinesisFirehoseDeliveryStreamEncryptionConfiguration struct {
    
    // KMSEncryptionConfig AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration-kmsencryptionconfig
    KMSEncryptionConfig AWSKinesisFirehoseDeliveryStreamEncryptionConfigurationKMSEncryptionConfig `json:"KMSEncryptionConfig"`
    
    // NoEncryptionConfig AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration-noencryptionconfig
    NoEncryptionConfig string `json:"NoEncryptionConfig"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStreamEncryptionConfiguration) AWSCloudFormationType() string {
    return "AWS::KinesisFirehose::DeliveryStream.EncryptionConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStreamEncryptionConfiguration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}