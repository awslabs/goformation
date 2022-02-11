package lex

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// BotVersion_BotVersionLocaleDetails AWS CloudFormation Resource (AWS::Lex::BotVersion.BotVersionLocaleDetails)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botversion-botversionlocaledetails.html
type BotVersion_BotVersionLocaleDetails struct {

	// SourceBotVersion AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botversion-botversionlocaledetails.html#cfn-lex-botversion-botversionlocaledetails-sourcebotversion
	SourceBotVersion string `json:"SourceBotVersion"`

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
func (r *BotVersion_BotVersionLocaleDetails) AWSCloudFormationType() string {
	return "AWS::Lex::BotVersion.BotVersionLocaleDetails"
}
