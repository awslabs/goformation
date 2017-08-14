package resources

// AWS::S3::Bucket AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html
type AWSS3Bucket struct {

	// AccessControl AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-accesscontrol

	AccessControl string `json:"AccessControl"`

	// BucketName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-name

	BucketName string `json:"BucketName"`

	// CorsConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-crossoriginconfig

	CorsConfiguration AWSS3Bucket_CorsConfiguration `json:"CorsConfiguration"`

	// LifecycleConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-lifecycleconfig

	LifecycleConfiguration AWSS3Bucket_LifecycleConfiguration `json:"LifecycleConfiguration"`

	// LoggingConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-loggingconfig

	LoggingConfiguration AWSS3Bucket_LoggingConfiguration `json:"LoggingConfiguration"`

	// NotificationConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-notification

	NotificationConfiguration AWSS3Bucket_NotificationConfiguration `json:"NotificationConfiguration"`

	// ReplicationConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-replicationconfiguration

	ReplicationConfiguration AWSS3Bucket_ReplicationConfiguration `json:"ReplicationConfiguration"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-tags

	Tags []Tag `json:"Tags"`

	// VersioningConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-versioning

	VersioningConfiguration AWSS3Bucket_VersioningConfiguration `json:"VersioningConfiguration"`

	// WebsiteConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-websiteconfiguration

	WebsiteConfiguration AWSS3Bucket_WebsiteConfiguration `json:"WebsiteConfiguration"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket) AWSCloudFormationType() string {
	return "AWS::S3::Bucket"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
