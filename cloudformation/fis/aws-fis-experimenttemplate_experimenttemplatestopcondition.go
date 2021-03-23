package fis

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// ExperimentTemplate_ExperimentTemplateStopCondition AWS CloudFormation Resource (AWS::FIS::ExperimentTemplate.ExperimentTemplateStopCondition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatestopcondition.html
type ExperimentTemplate_ExperimentTemplateStopCondition struct {

	// source AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatestopcondition.html#cfn-fis-experimenttemplate-experimenttemplatestopcondition-source
	source string `json:"source,omitempty"`

	// value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatestopcondition.html#cfn-fis-experimenttemplate-experimenttemplatestopcondition-value
	value string `json:"value,omitempty"`

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
func (r *ExperimentTemplate_ExperimentTemplateStopCondition) AWSCloudFormationType() string {
	return "AWS::FIS::ExperimentTemplate.ExperimentTemplateStopCondition"
}
