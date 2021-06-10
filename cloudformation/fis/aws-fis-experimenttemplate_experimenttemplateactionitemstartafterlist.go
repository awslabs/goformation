package fis

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// ExperimentTemplate_ExperimentTemplateActionItemStartAfterList AWS CloudFormation Resource (AWS::FIS::ExperimentTemplate.ExperimentTemplateActionItemStartAfterList)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateactionitemstartafterlist.html
type ExperimentTemplate_ExperimentTemplateActionItemStartAfterList struct {

	// ExperimentTemplateActionItemStartAfterList AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateactionitemstartafterlist.html#cfn-fis-experimenttemplate-experimenttemplateactionitemstartafterlist-experimenttemplateactionitemstartafterlist
	ExperimentTemplateActionItemStartAfterList []string `json:"ExperimentTemplateActionItemStartAfterList,omitempty"`

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
func (r *ExperimentTemplate_ExperimentTemplateActionItemStartAfterList) AWSCloudFormationType() string {
	return "AWS::FIS::ExperimentTemplate.ExperimentTemplateActionItemStartAfterList"
}
