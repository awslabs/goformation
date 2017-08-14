package resources

// AWS::S3::Bucket.LoggingConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfig.html
type AWSS3BucketLoggingConfiguration struct {

	// DestinationBucketName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfig.html#cfn-s3-bucket-loggingconfig-destinationbucketname
	DestinationBucketName string `json:"DestinationBucketName"`

	// LogFilePrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfig.html#cfn-s3-bucket-loggingconfig-logfileprefix
	LogFilePrefix string `json:"LogFilePrefix"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3BucketLoggingConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.LoggingConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3BucketLoggingConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
