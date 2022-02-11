package lex

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// BotAlias_BotAliasLocaleSettings AWS CloudFormation Resource (AWS::Lex::BotAlias.BotAliasLocaleSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-botaliaslocalesettings.html
type BotAlias_BotAliasLocaleSettings struct {

	// CodeHookSpecification AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-botaliaslocalesettings.html#cfn-lex-botalias-botaliaslocalesettings-codehookspecification
	CodeHookSpecification *BotAlias_CodeHookSpecification `json:"CodeHookSpecification,omitempty"`

	// Enabled AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-botaliaslocalesettings.html#cfn-lex-botalias-botaliaslocalesettings-enabled
	Enabled bool `json:"Enabled"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *BotAlias_BotAliasLocaleSettings) AWSCloudFormationType() string {
	return "AWS::Lex::BotAlias.BotAliasLocaleSettings"
}
