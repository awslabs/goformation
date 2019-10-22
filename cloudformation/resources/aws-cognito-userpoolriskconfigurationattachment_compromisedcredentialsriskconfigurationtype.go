package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSCognitoUserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType AWS CloudFormation Resource (AWS::Cognito::UserPoolRiskConfigurationAttachment.CompromisedCredentialsRiskConfigurationType)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype.html
type AWSCognitoUserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType struct {

	// Actions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype-actions
	Actions *AWSCognitoUserPoolRiskConfigurationAttachment_CompromisedCredentialsActionsType `json:"Actions,omitempty"`

	// EventFilter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype-eventfilter
	EventFilter []string `json:"EventFilter,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPoolRiskConfigurationAttachment.CompromisedCredentialsRiskConfigurationType"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSCognitoUserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSCognitoUserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSCognitoUserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSCognitoUserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSCognitoUserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
