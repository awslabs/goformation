package iotsitewise

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// Portal_PortalStatus AWS CloudFormation Resource (AWS::IoTSiteWise::Portal.PortalStatus)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-portal-portalstatus.html
type Portal_PortalStatus struct {

	// error AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-portal-portalstatus.html#cfn-iotsitewise-portal-portalstatus-error
	error *Portal_MonitorErrorDetails `json:"error,omitempty"`

	// state AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-portal-portalstatus.html#cfn-iotsitewise-portal-portalstatus-state
	state string `json:"state,omitempty"`

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
func (r *Portal_PortalStatus) AWSCloudFormationType() string {
	return "AWS::IoTSiteWise::Portal.PortalStatus"
}
