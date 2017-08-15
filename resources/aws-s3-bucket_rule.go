package resources

// AWS::S3::Bucket.Rule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html
type AWSS3Bucket_Rule struct {

	// ExpirationDate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-expirationdate
	ExpirationDate string `json:"ExpirationDate"`

	// ExpirationInDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-expirationindays
	ExpirationInDays int64 `json:"ExpirationInDays"`

	// Id AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-id
	Id string `json:"Id"`

	// NoncurrentVersionExpirationInDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversionexpirationindays
	NoncurrentVersionExpirationInDays int64 `json:"NoncurrentVersionExpirationInDays"`

	// NoncurrentVersionTransition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition
	NoncurrentVersionTransition AWSS3Bucket_NoncurrentVersionTransition `json:"NoncurrentVersionTransition"`

	// NoncurrentVersionTransitions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversiontransitions
	NoncurrentVersionTransitions AWSS3Bucket_NoncurrentVersionTransition `json:"NoncurrentVersionTransitions"`

	// Prefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-prefix
	Prefix string `json:"Prefix"`

	// Status AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-status
	Status string `json:"Status"`

	// Transition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-transition
	Transition AWSS3Bucket_Transition `json:"Transition"`

	// Transitions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-transitions
	Transitions AWSS3Bucket_Transition `json:"Transitions"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_Rule) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.Rule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_Rule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
