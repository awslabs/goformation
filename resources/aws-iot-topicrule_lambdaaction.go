package resources

// AWS::IoT::TopicRule.LambdaAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-lambda.html
type AWSIoTTopicRuleLambdaAction struct {
    
    // FunctionArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-lambda.html#cfn-iot-lambda-functionarn
    FunctionArn string `json:"FunctionArn"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRuleLambdaAction) AWSCloudFormationType() string {
    return "AWS::IoT::TopicRule.LambdaAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRuleLambdaAction) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}