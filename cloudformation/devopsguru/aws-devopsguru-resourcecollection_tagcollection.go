// Code generated by "go generate". Please don't change this file directly.

package devopsguru

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// ResourceCollection_TagCollection AWS CloudFormation Resource (AWS::DevOpsGuru::ResourceCollection.TagCollection)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-resourcecollection-tagcollection.html
type ResourceCollection_TagCollection struct {

	// AppBoundaryKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-resourcecollection-tagcollection.html#cfn-devopsguru-resourcecollection-tagcollection-appboundarykey
	AppBoundaryKey *string `json:"AppBoundaryKey,omitempty"`

	// TagValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-resourcecollection-tagcollection.html#cfn-devopsguru-resourcecollection-tagcollection-tagvalues
	TagValues []string `json:"TagValues,omitempty"`

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
func (r *ResourceCollection_TagCollection) AWSCloudFormationType() string {
	return "AWS::DevOpsGuru::ResourceCollection.TagCollection"
}