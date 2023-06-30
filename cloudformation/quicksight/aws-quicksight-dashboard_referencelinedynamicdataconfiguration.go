// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_ReferenceLineDynamicDataConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.ReferenceLineDynamicDataConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-referencelinedynamicdataconfiguration.html
type Dashboard_ReferenceLineDynamicDataConfiguration struct {

	// Calculation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-referencelinedynamicdataconfiguration.html#cfn-quicksight-dashboard-referencelinedynamicdataconfiguration-calculation
	Calculation *Dashboard_NumericalAggregationFunction `json:"Calculation"`

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-referencelinedynamicdataconfiguration.html#cfn-quicksight-dashboard-referencelinedynamicdataconfiguration-column
	Column *Dashboard_ColumnIdentifier `json:"Column"`

	// MeasureAggregationFunction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-referencelinedynamicdataconfiguration.html#cfn-quicksight-dashboard-referencelinedynamicdataconfiguration-measureaggregationfunction
	MeasureAggregationFunction *Dashboard_AggregationFunction `json:"MeasureAggregationFunction,omitempty"`

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
func (r *Dashboard_ReferenceLineDynamicDataConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.ReferenceLineDynamicDataConfiguration"
}
