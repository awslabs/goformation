package resources

// AWS::ECS::TaskDefinition.HostVolumeProperties AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes-host.html
type AWSECSTaskDefinitionHostVolumeProperties struct {

	// SourcePath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes-host.html#cfn-ecs-taskdefinition-volumes-host-sourcepath
	SourcePath string `json:"SourcePath"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinitionHostVolumeProperties) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.HostVolumeProperties"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinitionHostVolumeProperties) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
