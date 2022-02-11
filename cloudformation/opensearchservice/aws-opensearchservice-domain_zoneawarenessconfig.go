package opensearchservice

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// Domain_ZoneAwarenessConfig AWS CloudFormation Resource (AWS::OpenSearchService::Domain.ZoneAwarenessConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-zoneawarenessconfig.html
type Domain_ZoneAwarenessConfig struct {

	// AvailabilityZoneCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-zoneawarenessconfig.html#cfn-opensearchservice-domain-zoneawarenessconfig-availabilityzonecount
	AvailabilityZoneCount *int `json:"AvailabilityZoneCount,omitempty"`

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
func (r *Domain_ZoneAwarenessConfig) AWSCloudFormationType() string {
	return "AWS::OpenSearchService::Domain.ZoneAwarenessConfig"
}
