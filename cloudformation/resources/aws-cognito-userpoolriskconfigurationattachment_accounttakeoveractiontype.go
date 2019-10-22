package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSCognitoUserPoolRiskConfigurationAttachment_AccountTakeoverActionType AWS CloudFormation Resource (AWS::Cognito::UserPoolRiskConfigurationAttachment.AccountTakeoverActionType)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype.html
type AWSCognitoUserPoolRiskConfigurationAttachment_AccountTakeoverActionType struct {

	// EventAction AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype-eventaction
	EventAction string `json:"EventAction,omitempty"`

	// Notify AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype-notify
	Notify bool `json:"Notify"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolRiskConfigurationAttachment_AccountTakeoverActionType) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPoolRiskConfigurationAttachment.AccountTakeoverActionType"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSCognitoUserPoolRiskConfigurationAttachment_AccountTakeoverActionType) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSCognitoUserPoolRiskConfigurationAttachment_AccountTakeoverActionType) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSCognitoUserPoolRiskConfigurationAttachment_AccountTakeoverActionType) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSCognitoUserPoolRiskConfigurationAttachment_AccountTakeoverActionType) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSCognitoUserPoolRiskConfigurationAttachment_AccountTakeoverActionType) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
