package cloudformation

// AWSECSTaskDefinition_DockerVolumeConfiguration AWS CloudFormation Resource (AWS::ECS::TaskDefinition.DockerVolumeConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-dockervolumeconfiguration.html
type AWSECSTaskDefinition_DockerVolumeConfiguration struct {

	// Autoprovision AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-dockervolumeconfiguration.html#cfn-ecs-taskdefinition-dockervolumeconfiguration-autoprovision
	Autoprovision bool `json:"Autoprovision,omitempty"`

	// Driver AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-dockervolumeconfiguration.html#cfn-ecs-taskdefinition-dockervolumeconfiguration-driver
	Driver string `json:"Driver,omitempty"`

	// DriverOpts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-dockervolumeconfiguration.html#cfn-ecs-taskdefinition-dockervolumeconfiguration-driveropts
	DriverOpts map[string]string `json:"DriverOpts,omitempty"`

	// Labels AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-dockervolumeconfiguration.html#cfn-ecs-taskdefinition-dockervolumeconfiguration-labels
	Labels map[string]string `json:"Labels,omitempty"`

	// Scope AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-dockervolumeconfiguration.html#cfn-ecs-taskdefinition-dockervolumeconfiguration-scope
	Scope string `json:"Scope,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_DockerVolumeConfiguration) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.DockerVolumeConfiguration"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSECSTaskDefinition_DockerVolumeConfiguration) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
