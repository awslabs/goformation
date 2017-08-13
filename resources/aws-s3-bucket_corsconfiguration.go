package resources

// AWS::S3::Bucket.CorsConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors.html
type AWSS3BucketCorsConfiguration struct {
    
    // CorsRules AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors.html#cfn-s3-bucket-cors-corsrule
    CorsRules []AWSS3BucketCorsConfigurationCorsRule `json:"CorsRules"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3BucketCorsConfiguration) AWSCloudFormationType() string {
    return "AWS::S3::Bucket.CorsConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3BucketCorsConfiguration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}