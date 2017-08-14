package resources

// AWS::S3::Bucket.QueueConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html
type AWSS3BucketQueueConfiguration struct {

	// Event AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html#cfn-s3-bucket-notificationconfig-queueconfig-event
	Event string `json:"Event"`

	// Filter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html#cfn-s3-bucket-notificationconfig-queueconfig-filter
	Filter AWSS3BucketQueueConfigurationNotificationFilter `json:"Filter"`

	// Queue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html#cfn-s3-bucket-notificationconfig-queueconfig-queue
	Queue string `json:"Queue"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3BucketQueueConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.QueueConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3BucketQueueConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
