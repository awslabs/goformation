package resources

// AWS::AutoScaling::ScheduledAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html
type AWSAutoScalingScheduledAction struct {
    
    // AutoScalingGroupName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-asgname
    AutoScalingGroupName string `json:"AutoScalingGroupName"`
    
    // DesiredCapacity AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-desiredcapacity
    DesiredCapacity int64 `json:"DesiredCapacity"`
    
    // EndTime AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-endtime
    EndTime string `json:"EndTime"`
    
    // MaxSize AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-maxsize
    MaxSize int64 `json:"MaxSize"`
    
    // MinSize AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-minsize
    MinSize int64 `json:"MinSize"`
    
    // Recurrence AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-recurrence
    Recurrence string `json:"Recurrence"`
    
    // StartTime AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-starttime
    StartTime string `json:"StartTime"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingScheduledAction) AWSCloudFormationType() string {
    return "AWS::AutoScaling::ScheduledAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSAutoScalingScheduledAction) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}