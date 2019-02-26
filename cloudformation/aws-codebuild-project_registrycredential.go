package cloudformation

// AWSCodeBuildProject_RegistryCredential AWS CloudFormation Resource (AWS::CodeBuild::Project.RegistryCredential)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-registrycredential.html
type AWSCodeBuildProject_RegistryCredential struct {

	// Credential AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-registrycredential.html#cfn-codebuild-project-registrycredential-credential
	Credential string `json:"Credential,omitempty"`

	// CredentialProvider AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-registrycredential.html#cfn-codebuild-project-registrycredential-credentialprovider
	CredentialProvider string `json:"CredentialProvider,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeBuildProject_RegistryCredential) AWSCloudFormationType() string {
	return "AWS::CodeBuild::Project.RegistryCredential"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSCodeBuildProject_RegistryCredential) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
