package lex

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// Bot_GrammarSlotTypeSource AWS CloudFormation Resource (AWS::Lex::Bot.GrammarSlotTypeSource)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-grammarslottypesource.html
type Bot_GrammarSlotTypeSource struct {

	// KmsKeyArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-grammarslottypesource.html#cfn-lex-bot-grammarslottypesource-kmskeyarn
	KmsKeyArn string `json:"KmsKeyArn,omitempty"`

	// S3BucketName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-grammarslottypesource.html#cfn-lex-bot-grammarslottypesource-s3bucketname
	S3BucketName string `json:"S3BucketName,omitempty"`

	// S3ObjectKey AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-grammarslottypesource.html#cfn-lex-bot-grammarslottypesource-s3objectkey
	S3ObjectKey string `json:"S3ObjectKey,omitempty"`

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
func (r *Bot_GrammarSlotTypeSource) AWSCloudFormationType() string {
	return "AWS::Lex::Bot.GrammarSlotTypeSource"
}