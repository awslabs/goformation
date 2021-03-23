package fis

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// ExperimentTemplate_ExperimentTemplateAction AWS CloudFormation Resource (AWS::FIS::ExperimentTemplate.ExperimentTemplateAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html
type ExperimentTemplate_ExperimentTemplateAction struct {

	// actionId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html#cfn-fis-experimenttemplate-experimenttemplateaction-actionid
	actionId string `json:"actionId,omitempty"`

	// description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html#cfn-fis-experimenttemplate-experimenttemplateaction-description
	description string `json:"description,omitempty"`

	// parameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html#cfn-fis-experimenttemplate-experimenttemplateaction-parameters
	parameters *ExperimentTemplate_ExperimentTemplateActionItemParameterMap `json:"parameters,omitempty"`

	// startAfter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html#cfn-fis-experimenttemplate-experimenttemplateaction-startafter
	startAfter *ExperimentTemplate_ExperimentTemplateActionItemStartAfterList `json:"startAfter,omitempty"`

	// targets AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html#cfn-fis-experimenttemplate-experimenttemplateaction-targets
	targets *ExperimentTemplate_ExperimentTemplateActionItemTargetMap `json:"targets,omitempty"`

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
func (r *ExperimentTemplate_ExperimentTemplateAction) AWSCloudFormationType() string {
	return "AWS::FIS::ExperimentTemplate.ExperimentTemplateAction"
}
