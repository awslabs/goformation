package resources

// AWS::ECS::TaskDefinition.HostEntry AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html
type AWSECSTaskDefinitionHostEntry struct {
    
    // Hostname AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html#cfn-ecs-taskdefinition-containerdefinition-hostentry-hostname
    Hostname string `json:"Hostname"`
    
    // IpAddress AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html#cfn-ecs-taskdefinition-containerdefinition-hostentry-ipaddress
    IpAddress string `json:"IpAddress"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinitionHostEntry) AWSCloudFormationType() string {
    return "AWS::ECS::TaskDefinition.HostEntry"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinitionHostEntry) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}