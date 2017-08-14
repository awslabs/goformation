package resources

// AWS::S3::Bucket.NotificationFilter AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter.html
type AWSS3BucketNotificationFilter struct {

	// S3Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter.html#cfn-s3-bucket-notificationconfiguraiton-config-filter-s3key
	S3Key AWSS3BucketNotificationFilterS3KeyFilter `json:"S3Key"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3BucketNotificationFilter) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.NotificationFilter"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3BucketNotificationFilter) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
