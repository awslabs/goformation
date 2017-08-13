package resources

// AWS::EC2::FlowLog AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html
type AWSEC2FlowLog struct {
    
    // DeliverLogsPermissionArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-deliverlogspermissionarn
    DeliverLogsPermissionArn string `json:"DeliverLogsPermissionArn"`
    
    // LogGroupName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-loggroupname
    LogGroupName string `json:"LogGroupName"`
    
    // ResourceId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-resourceid
    ResourceId string `json:"ResourceId"`
    
    // ResourceType AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-resourcetype
    ResourceType string `json:"ResourceType"`
    
    // TrafficType AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-traffictype
    TrafficType string `json:"TrafficType"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2FlowLog) AWSCloudFormationType() string {
    return "AWS::EC2::FlowLog"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2FlowLog) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}