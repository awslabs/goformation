package resources

// AWS::ECS::TaskDefinition.Volume AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes.html
type AWSECSTaskDefinitionVolume struct {
    
    // Host AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes.html#cfn-ecs-taskdefinition-volumes-host
    Host AWSECSTaskDefinitionVolumeHostVolumeProperties `json:"Host"`
    
    // Name AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes.html#cfn-ecs-taskdefinition-volumes-name
    Name string `json:"Name"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinitionVolume) AWSCloudFormationType() string {
    return "AWS::ECS::TaskDefinition.Volume"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinitionVolume) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}