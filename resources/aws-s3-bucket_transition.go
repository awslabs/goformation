package resources

// AWS::S3::Bucket.Transition AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html
type AWSS3BucketTransition struct {

	// StorageClass AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html#cfn-s3-bucket-lifecycleconfig-rule-transition-storageclass
	StorageClass string `json:"StorageClass"`

	// TransitionDate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html#cfn-s3-bucket-lifecycleconfig-rule-transition-transitiondate
	TransitionDate time.Time `json:"TransitionDate"`

	// TransitionInDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html#cfn-s3-bucket-lifecycleconfig-rule-transition-transitionindays
	TransitionInDays int64 `json:"TransitionInDays"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3BucketTransition) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.Transition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3BucketTransition) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
