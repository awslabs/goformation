// Code generated by "go generate". Please don't change this file directly.

package comprehend

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Flywheel_TaskConfig AWS CloudFormation Resource (AWS::Comprehend::Flywheel.TaskConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-taskconfig.html
type Flywheel_TaskConfig struct {

	// DocumentClassificationConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-taskconfig.html#cfn-comprehend-flywheel-taskconfig-documentclassificationconfig
	DocumentClassificationConfig *Flywheel_DocumentClassificationConfig `json:"DocumentClassificationConfig,omitempty"`

	// EntityRecognitionConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-taskconfig.html#cfn-comprehend-flywheel-taskconfig-entityrecognitionconfig
	EntityRecognitionConfig *Flywheel_EntityRecognitionConfig `json:"EntityRecognitionConfig,omitempty"`

	// LanguageCode AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-taskconfig.html#cfn-comprehend-flywheel-taskconfig-languagecode
	LanguageCode string `json:"LanguageCode"`

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
func (r *Flywheel_TaskConfig) AWSCloudFormationType() string {
	return "AWS::Comprehend::Flywheel.TaskConfig"
}
