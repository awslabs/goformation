package resources

// AWS::S3::Bucket.NoncurrentVersionTransition AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition.html
type AWSS3BucketNoncurrentVersionTransition struct {
    
    // StorageClass AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition-storageclass
    StorageClass string `json:"StorageClass"`
    
    // TransitionInDays AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition-transitionindays
    TransitionInDays int64 `json:"TransitionInDays"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3BucketNoncurrentVersionTransition) AWSCloudFormationType() string {
    return "AWS::S3::Bucket.NoncurrentVersionTransition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3BucketNoncurrentVersionTransition) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}