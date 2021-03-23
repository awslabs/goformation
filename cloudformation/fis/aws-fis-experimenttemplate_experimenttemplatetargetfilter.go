package fis

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// ExperimentTemplate_ExperimentTemplateTargetFilter AWS CloudFormation Resource (AWS::FIS::ExperimentTemplate.ExperimentTemplateTargetFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetargetfilter.html
type ExperimentTemplate_ExperimentTemplateTargetFilter struct {

	// path AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetargetfilter.html#cfn-fis-experimenttemplate-experimenttemplatetargetfilter-path
	path string `json:"path,omitempty"`

	// values AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetargetfilter.html#cfn-fis-experimenttemplate-experimenttemplatetargetfilter-values
	values *ExperimentTemplate_ExperimentTemplateTargetFilterValues `json:"values,omitempty"`

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
func (r *ExperimentTemplate_ExperimentTemplateTargetFilter) AWSCloudFormationType() string {
	return "AWS::FIS::ExperimentTemplate.ExperimentTemplateTargetFilter"
}
