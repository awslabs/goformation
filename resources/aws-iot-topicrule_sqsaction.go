package resources

// AWS::IoT::TopicRule.SqsAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sqs.html
type AWSIoTTopicRuleSqsAction struct {
    
    // QueueUrl AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sqs.html#cfn-iot-sqs-queueurl
    QueueUrl string `json:"QueueUrl"`
    
    // RoleArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sqs.html#cfn-iot-sqs-rolearn
    RoleArn string `json:"RoleArn"`
    
    // UseBase64 AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sqs.html#cfn-iot-sqs-usebase64
    UseBase64 bool `json:"UseBase64"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRuleSqsAction) AWSCloudFormationType() string {
    return "AWS::IoT::TopicRule.SqsAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRuleSqsAction) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}