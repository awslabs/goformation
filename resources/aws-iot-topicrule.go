package resources

// AWS::IoT::TopicRule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html
type AWSIoTTopicRule struct {
    
    // RuleName AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html#cfn-iot-topicrule-rulename
    RuleName string `json:"RuleName"`
    
    // TopicRulePayload AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html#cfn-iot-topicrule-topicrulepayload
    TopicRulePayload AWSIoTTopicRuleTopicRulePayload `json:"TopicRulePayload"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule) AWSCloudFormationType() string {
    return "AWS::IoT::TopicRule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}