// Code generated by "go generate". Please don't change this file directly.

package batch

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// JobDefinition_PodProperties AWS CloudFormation Resource (AWS::Batch::JobDefinition.PodProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html
type JobDefinition_PodProperties struct {

	// Containers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html#cfn-batch-jobdefinition-podproperties-containers
	Containers []JobDefinition_EksContainer `json:"Containers,omitempty"`

	// DnsPolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html#cfn-batch-jobdefinition-podproperties-dnspolicy
	DnsPolicy *string `json:"DnsPolicy,omitempty"`

	// HostNetwork AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html#cfn-batch-jobdefinition-podproperties-hostnetwork
	HostNetwork *bool `json:"HostNetwork,omitempty"`

	// Metadata AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html#cfn-batch-jobdefinition-podproperties-metadata
	Metadata *JobDefinition_Metadata `json:"Metadata,omitempty"`

	// ServiceAccountName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html#cfn-batch-jobdefinition-podproperties-serviceaccountname
	ServiceAccountName *string `json:"ServiceAccountName,omitempty"`

	// Volumes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html#cfn-batch-jobdefinition-podproperties-volumes
	Volumes []JobDefinition_EksVolume `json:"Volumes,omitempty"`

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
func (r *JobDefinition_PodProperties) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.PodProperties"
}
