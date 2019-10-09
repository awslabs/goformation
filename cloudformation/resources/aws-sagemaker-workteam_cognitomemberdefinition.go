package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSSageMakerWorkteam_CognitoMemberDefinition AWS CloudFormation Resource (AWS::SageMaker::Workteam.CognitoMemberDefinition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-cognitomemberdefinition.html
type AWSSageMakerWorkteam_CognitoMemberDefinition struct {

	// CognitoClientId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-cognitomemberdefinition.html#cfn-sagemaker-workteam-cognitomemberdefinition-cognitoclientid
	CognitoClientId string `json:"CognitoClientId,omitempty"`

	// CognitoUserGroup AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-cognitomemberdefinition.html#cfn-sagemaker-workteam-cognitomemberdefinition-cognitousergroup
	CognitoUserGroup string `json:"CognitoUserGroup,omitempty"`

	// CognitoUserPool AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-cognitomemberdefinition.html#cfn-sagemaker-workteam-cognitomemberdefinition-cognitouserpool
	CognitoUserPool string `json:"CognitoUserPool,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSageMakerWorkteam_CognitoMemberDefinition) AWSCloudFormationType() string {
	return "AWS::SageMaker::Workteam.CognitoMemberDefinition"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSSageMakerWorkteam_CognitoMemberDefinition) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSSageMakerWorkteam_CognitoMemberDefinition) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSSageMakerWorkteam_CognitoMemberDefinition) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSSageMakerWorkteam_CognitoMemberDefinition) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSSageMakerWorkteam_CognitoMemberDefinition) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSSageMakerWorkteam_CognitoMemberDefinition) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
