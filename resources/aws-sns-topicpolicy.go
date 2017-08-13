package resources

// AWS::SNS::TopicPolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html
type AWSSNSTopicPolicy struct {
    
    // PolicyDocument AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html#cfn-sns-topicpolicy-policydocument
    PolicyDocument string `json:"PolicyDocument"`
    
    // Topics AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html#cfn-sns-topicpolicy-topics
    Topics []string `json:"Topics"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSNSTopicPolicy) AWSCloudFormationType() string {
    return "AWS::SNS::TopicPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSNSTopicPolicy) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}