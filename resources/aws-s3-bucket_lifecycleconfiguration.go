package resources

// AWS::S3::Bucket.LifecycleConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig.html
type AWSS3BucketLifecycleConfiguration struct {
    
    // Rules AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig.html#cfn-s3-bucket-lifecycleconfig-rules
    Rules []AWSS3BucketLifecycleConfigurationRule `json:"Rules"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3BucketLifecycleConfiguration) AWSCloudFormationType() string {
    return "AWS::S3::Bucket.LifecycleConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3BucketLifecycleConfiguration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}