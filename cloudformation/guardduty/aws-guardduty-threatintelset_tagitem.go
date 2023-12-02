// Code generated by "go generate". Please don't change this file directly.

package guardduty

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// ThreatIntelSet_TagItem AWS CloudFormation Resource (AWS::GuardDuty::ThreatIntelSet.TagItem)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-threatintelset-tagitem.html
type ThreatIntelSet_TagItem struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-threatintelset-tagitem.html#cfn-guardduty-threatintelset-tagitem-key
	Key string `json:"Key"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-threatintelset-tagitem.html#cfn-guardduty-threatintelset-tagitem-value
	Value string `json:"Value"`

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
func (r *ThreatIntelSet_TagItem) AWSCloudFormationType() string {
	return "AWS::GuardDuty::ThreatIntelSet.TagItem"
}
