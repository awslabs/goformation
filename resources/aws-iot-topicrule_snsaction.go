package resources

// AWS::IoT::TopicRule.SnsAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sns.html
type AWSIoTTopicRuleSnsAction struct {
    
    // MessageFormat AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sns.html#cfn-iot-sns-snsaction
    MessageFormat string `json:"MessageFormat"`
    
    // RoleArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sns.html#cfn-iot-sns-rolearn
    RoleArn string `json:"RoleArn"`
    
    // TargetArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sns.html#cfn-iot-sns-targetarn
    TargetArn string `json:"TargetArn"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRuleSnsAction) AWSCloudFormationType() string {
    return "AWS::IoT::TopicRule.SnsAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRuleSnsAction) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}