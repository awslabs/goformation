// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_FormatConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.FormatConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-formatconfiguration.html
type Analysis_FormatConfiguration struct {

	// DateTimeFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-formatconfiguration.html#cfn-quicksight-analysis-formatconfiguration-datetimeformatconfiguration
	DateTimeFormatConfiguration *Analysis_DateTimeFormatConfiguration `json:"DateTimeFormatConfiguration,omitempty"`

	// NumberFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-formatconfiguration.html#cfn-quicksight-analysis-formatconfiguration-numberformatconfiguration
	NumberFormatConfiguration *Analysis_NumberFormatConfiguration `json:"NumberFormatConfiguration,omitempty"`

	// StringFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-formatconfiguration.html#cfn-quicksight-analysis-formatconfiguration-stringformatconfiguration
	StringFormatConfiguration *Analysis_StringFormatConfiguration `json:"StringFormatConfiguration,omitempty"`

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
func (r *Analysis_FormatConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.FormatConfiguration"
}
