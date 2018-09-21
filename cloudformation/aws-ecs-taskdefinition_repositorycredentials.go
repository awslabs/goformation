package cloudformation

// AWSECSTaskDefinition_RepositoryCredentials AWS CloudFormation Resource (AWS::ECS::TaskDefinition.RepositoryCredentials)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-repositorycredentials.html
type AWSECSTaskDefinition_RepositoryCredentials struct {

	// CredentialsParameter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-repositorycredentials.html#cfn-ecs-taskdefinition-repositorycredentials-credentialsparameter
	CredentialsParameter string `json:"CredentialsParameter,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_RepositoryCredentials) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.RepositoryCredentials"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSECSTaskDefinition_RepositoryCredentials) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
