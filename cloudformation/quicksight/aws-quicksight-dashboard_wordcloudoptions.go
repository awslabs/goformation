// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_WordCloudOptions AWS CloudFormation Resource (AWS::QuickSight::Dashboard.WordCloudOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-wordcloudoptions.html
type Dashboard_WordCloudOptions struct {

	// CloudLayout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-wordcloudoptions.html#cfn-quicksight-dashboard-wordcloudoptions-cloudlayout
	CloudLayout *string `json:"CloudLayout,omitempty"`

	// MaximumStringLength AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-wordcloudoptions.html#cfn-quicksight-dashboard-wordcloudoptions-maximumstringlength
	MaximumStringLength *float64 `json:"MaximumStringLength,omitempty"`

	// WordCasing AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-wordcloudoptions.html#cfn-quicksight-dashboard-wordcloudoptions-wordcasing
	WordCasing *string `json:"WordCasing,omitempty"`

	// WordOrientation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-wordcloudoptions.html#cfn-quicksight-dashboard-wordcloudoptions-wordorientation
	WordOrientation *string `json:"WordOrientation,omitempty"`

	// WordPadding AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-wordcloudoptions.html#cfn-quicksight-dashboard-wordcloudoptions-wordpadding
	WordPadding *string `json:"WordPadding,omitempty"`

	// WordScaling AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-wordcloudoptions.html#cfn-quicksight-dashboard-wordcloudoptions-wordscaling
	WordScaling *string `json:"WordScaling,omitempty"`

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
func (r *Dashboard_WordCloudOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.WordCloudOptions"
}