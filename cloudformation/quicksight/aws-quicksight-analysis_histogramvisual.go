// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_HistogramVisual AWS CloudFormation Resource (AWS::QuickSight::Analysis.HistogramVisual)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramvisual.html
type Analysis_HistogramVisual struct {

	// Actions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramvisual.html#cfn-quicksight-analysis-histogramvisual-actions
	Actions []Analysis_VisualCustomAction `json:"Actions,omitempty"`

	// ChartConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramvisual.html#cfn-quicksight-analysis-histogramvisual-chartconfiguration
	ChartConfiguration *Analysis_HistogramConfiguration `json:"ChartConfiguration,omitempty"`

	// Subtitle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramvisual.html#cfn-quicksight-analysis-histogramvisual-subtitle
	Subtitle *Analysis_VisualSubtitleLabelOptions `json:"Subtitle,omitempty"`

	// Title AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramvisual.html#cfn-quicksight-analysis-histogramvisual-title
	Title *Analysis_VisualTitleLabelOptions `json:"Title,omitempty"`

	// VisualId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramvisual.html#cfn-quicksight-analysis-histogramvisual-visualid
	VisualId string `json:"VisualId"`

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
func (r *Analysis_HistogramVisual) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.HistogramVisual"
}