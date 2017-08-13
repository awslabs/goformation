package resources

// AWS::ECS::TaskDefinition.MountPoint AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html
type AWSECSTaskDefinitionMountPoint struct {
    
    // ContainerPath AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html#cfn-ecs-taskdefinition-containerdefinition-mountpoints-containerpath
    ContainerPath string `json:"ContainerPath"`
    
    // ReadOnly AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html#cfn-ecs-taskdefinition-containerdefinition-mountpoints-readonly
    ReadOnly bool `json:"ReadOnly"`
    
    // SourceVolume AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html#cfn-ecs-taskdefinition-containerdefinition-mountpoints-sourcevolume
    SourceVolume string `json:"SourceVolume"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinitionMountPoint) AWSCloudFormationType() string {
    return "AWS::ECS::TaskDefinition.MountPoint"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinitionMountPoint) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}