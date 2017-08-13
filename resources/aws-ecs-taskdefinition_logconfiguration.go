package resources

// AWS::ECS::TaskDefinition.LogConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-logconfiguration.html
type AWSECSTaskDefinitionLogConfiguration struct {
    
    // LogDriver AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-logconfiguration.html#cfn-ecs-taskdefinition-containerdefinition-logconfiguration-logdriver
    LogDriver string `json:"LogDriver"`
    
    // Options AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-logconfiguration.html#cfn-ecs-taskdefinition-containerdefinition-logconfiguration-options
    Options map[string]string `json:"Options"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinitionLogConfiguration) AWSCloudFormationType() string {
    return "AWS::ECS::TaskDefinition.LogConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinitionLogConfiguration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}