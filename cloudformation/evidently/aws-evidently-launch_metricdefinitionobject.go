package evidently

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// Launch_MetricDefinitionObject AWS CloudFormation Resource (AWS::Evidently::Launch.MetricDefinitionObject)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-metricdefinitionobject.html
type Launch_MetricDefinitionObject struct {

	// EntityIdKey AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-metricdefinitionobject.html#cfn-evidently-launch-metricdefinitionobject-entityidkey
	EntityIdKey string `json:"EntityIdKey"`

	// EventPattern AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-metricdefinitionobject.html#cfn-evidently-launch-metricdefinitionobject-eventpattern
	EventPattern string `json:"EventPattern"`

	// MetricName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-metricdefinitionobject.html#cfn-evidently-launch-metricdefinitionobject-metricname
	MetricName string `json:"MetricName"`

	// UnitLabel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-metricdefinitionobject.html#cfn-evidently-launch-metricdefinitionobject-unitlabel
	UnitLabel *string `json:"UnitLabel,omitempty"`

	// ValueKey AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-metricdefinitionobject.html#cfn-evidently-launch-metricdefinitionobject-valuekey
	ValueKey string `json:"ValueKey"`

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
func (r *Launch_MetricDefinitionObject) AWSCloudFormationType() string {
	return "AWS::Evidently::Launch.MetricDefinitionObject"
}
