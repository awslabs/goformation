package lex

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Bot_Intent AWS CloudFormation Resource (AWS::Lex::Bot.Intent)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html
type Bot_Intent struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-description
	Description *string `json:"Description,omitempty"`

	// DialogCodeHook AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-dialogcodehook
	DialogCodeHook *Bot_DialogCodeHookSetting `json:"DialogCodeHook,omitempty"`

	// FulfillmentCodeHook AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-fulfillmentcodehook
	FulfillmentCodeHook *Bot_FulfillmentCodeHookSetting `json:"FulfillmentCodeHook,omitempty"`

	// InputContexts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-inputcontexts
	InputContexts *[]Bot_InputContext `json:"InputContexts,omitempty"`

	// IntentClosingSetting AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-intentclosingsetting
	IntentClosingSetting *Bot_IntentClosingSetting `json:"IntentClosingSetting,omitempty"`

	// IntentConfirmationSetting AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-intentconfirmationsetting
	IntentConfirmationSetting *Bot_IntentConfirmationSetting `json:"IntentConfirmationSetting,omitempty"`

	// KendraConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-kendraconfiguration
	KendraConfiguration *Bot_KendraConfiguration `json:"KendraConfiguration,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-name
	Name string `json:"Name"`

	// OutputContexts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-outputcontexts
	OutputContexts *[]Bot_OutputContext `json:"OutputContexts,omitempty"`

	// ParentIntentSignature AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-parentintentsignature
	ParentIntentSignature *string `json:"ParentIntentSignature,omitempty"`

	// SampleUtterances AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-sampleutterances
	SampleUtterances *[]Bot_SampleUtterance `json:"SampleUtterances,omitempty"`

	// SlotPriorities AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-slotpriorities
	SlotPriorities *[]Bot_SlotPriority `json:"SlotPriorities,omitempty"`

	// Slots AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-slots
	Slots *[]Bot_Slot `json:"Slots,omitempty"`

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
func (r *Bot_Intent) AWSCloudFormationType() string {
	return "AWS::Lex::Bot.Intent"
}
