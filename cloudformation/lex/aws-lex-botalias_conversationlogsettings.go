package lex

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// BotAlias_ConversationLogSettings AWS CloudFormation Resource (AWS::Lex::BotAlias.ConversationLogSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-conversationlogsettings.html
type BotAlias_ConversationLogSettings struct {

	// AudioLogSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-conversationlogsettings.html#cfn-lex-botalias-conversationlogsettings-audiologsettings
	AudioLogSettings *[]BotAlias_AudioLogSetting `json:"AudioLogSettings,omitempty"`

	// TextLogSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-conversationlogsettings.html#cfn-lex-botalias-conversationlogsettings-textlogsettings
	TextLogSettings *[]BotAlias_TextLogSetting `json:"TextLogSettings,omitempty"`

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
func (r *BotAlias_ConversationLogSettings) AWSCloudFormationType() string {
	return "AWS::Lex::BotAlias.ConversationLogSettings"
}
