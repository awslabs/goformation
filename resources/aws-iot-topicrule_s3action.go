package resources

// AWS::IoT::TopicRule.S3Action AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-s3.html
type AWSIoTTopicRuleS3Action struct {
    
    // BucketName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-s3.html#cfn-iot-s3-bucketname
    BucketName string `json:"BucketName"`
    
    // Key AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-s3.html#cfn-iot-s3-key
    Key string `json:"Key"`
    
    // RoleArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-s3.html#cfn-iot-s3-rolearn
    RoleArn string `json:"RoleArn"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRuleS3Action) AWSCloudFormationType() string {
    return "AWS::IoT::TopicRule.S3Action"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRuleS3Action) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}