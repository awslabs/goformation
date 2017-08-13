package resources

// AWS::S3::BucketPolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html
type AWSS3BucketPolicy struct {
    
    // Bucket AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html#cfn-s3-bucketpolicy-bucket
    Bucket string `json:"Bucket"`
    
    // PolicyDocument AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html#cfn-s3-bucketpolicy-policydocument
    PolicyDocument string `json:"PolicyDocument"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3BucketPolicy) AWSCloudFormationType() string {
    return "AWS::S3::BucketPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3BucketPolicy) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}