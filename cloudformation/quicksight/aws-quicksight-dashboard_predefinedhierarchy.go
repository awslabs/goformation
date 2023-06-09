// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_PredefinedHierarchy AWS CloudFormation Resource (AWS::QuickSight::Dashboard.PredefinedHierarchy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-predefinedhierarchy.html
type Dashboard_PredefinedHierarchy struct {

	// Columns AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-predefinedhierarchy.html#cfn-quicksight-dashboard-predefinedhierarchy-columns
	Columns []Dashboard_ColumnIdentifier `json:"Columns"`

	// DrillDownFilters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-predefinedhierarchy.html#cfn-quicksight-dashboard-predefinedhierarchy-drilldownfilters
	DrillDownFilters []Dashboard_DrillDownFilter `json:"DrillDownFilters,omitempty"`

	// HierarchyId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-predefinedhierarchy.html#cfn-quicksight-dashboard-predefinedhierarchy-hierarchyid
	HierarchyId string `json:"HierarchyId"`

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
func (r *Dashboard_PredefinedHierarchy) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.PredefinedHierarchy"
}