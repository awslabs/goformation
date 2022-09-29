// Code generated by "go generate". Please don't change this file directly.

package evidently

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Experiment_RunningStatusObject AWS CloudFormation Resource (AWS::Evidently::Experiment.RunningStatusObject)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-runningstatusobject.html
type Experiment_RunningStatusObject struct {

	// AnalysisCompleteTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-runningstatusobject.html#cfn-evidently-experiment-runningstatusobject-analysiscompletetime
	AnalysisCompleteTime *string `json:"AnalysisCompleteTime,omitempty"`

	// DesiredState AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-runningstatusobject.html#cfn-evidently-experiment-runningstatusobject-desiredstate
	DesiredState *string `json:"DesiredState,omitempty"`

	// Reason AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-runningstatusobject.html#cfn-evidently-experiment-runningstatusobject-reason
	Reason *string `json:"Reason,omitempty"`

	// Status AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-runningstatusobject.html#cfn-evidently-experiment-runningstatusobject-status
	Status *string `json:"Status,omitempty"`

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
func (r *Experiment_RunningStatusObject) AWSCloudFormationType() string {
	return "AWS::Evidently::Experiment.RunningStatusObject"
}
