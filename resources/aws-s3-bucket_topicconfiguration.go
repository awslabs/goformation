package resources

// AWS::S3::Bucket.TopicConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html
type AWSS3Bucket_TopicConfiguration struct {

	// Event AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html#cfn-s3-bucket-notificationconfig-topicconfig-event
	Event string `json:"Event"`

	// Filter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html#cfn-s3-bucket-notificationconfig-topicconfig-filter
	Filter AWSS3Bucket_NotificationFilter `json:"Filter"`

	// Topic AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html#cfn-s3-bucket-notificationconfig-topicconfig-topic
	Topic string `json:"Topic"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_TopicConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.TopicConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_TopicConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
