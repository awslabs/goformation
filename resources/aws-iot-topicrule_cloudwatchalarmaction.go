package resources

// AWS::IoT::TopicRule.CloudwatchAlarmAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchalarm.html
type AWSIoTTopicRuleCloudwatchAlarmAction struct {
    
    // AlarmName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchalarm.html#cfn-iot-cloudwatchalarm-alarmname
    AlarmName string `json:"AlarmName"`
    
    // RoleArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchalarm.html#cfn-iot-cloudwatchalarm-rolearn
    RoleArn string `json:"RoleArn"`
    
    // StateReason AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchalarm.html#cfn-iot-cloudwatchalarm-statereason
    StateReason string `json:"StateReason"`
    
    // StateValue AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchalarm.html#cfn-iot-cloudwatchalarm-statevalue
    StateValue string `json:"StateValue"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRuleCloudwatchAlarmAction) AWSCloudFormationType() string {
    return "AWS::IoT::TopicRule.CloudwatchAlarmAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRuleCloudwatchAlarmAction) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}