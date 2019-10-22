package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSSageMakerWorkteam_MemberDefinition AWS CloudFormation Resource (AWS::SageMaker::Workteam.MemberDefinition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-memberdefinition.html
type AWSSageMakerWorkteam_MemberDefinition struct {

	// CognitoMemberDefinition AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-memberdefinition.html#cfn-sagemaker-workteam-memberdefinition-cognitomemberdefinition
	CognitoMemberDefinition *AWSSageMakerWorkteam_CognitoMemberDefinition `json:"CognitoMemberDefinition,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSageMakerWorkteam_MemberDefinition) AWSCloudFormationType() string {
	return "AWS::SageMaker::Workteam.MemberDefinition"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSSageMakerWorkteam_MemberDefinition) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSSageMakerWorkteam_MemberDefinition) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSSageMakerWorkteam_MemberDefinition) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSSageMakerWorkteam_MemberDefinition) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSSageMakerWorkteam_MemberDefinition) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
