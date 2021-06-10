package wafv2

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// WebACL_CustomResponseBody AWS CloudFormation Resource (AWS::WAFv2::WebACL.CustomResponseBody)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-customresponsebody.html
type WebACL_CustomResponseBody struct {

	// Content AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-customresponsebody.html#cfn-wafv2-webacl-customresponsebody-content
	Content string `json:"Content,omitempty"`

	// ContentType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-customresponsebody.html#cfn-wafv2-webacl-customresponsebody-contenttype
	ContentType string `json:"ContentType,omitempty"`

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
func (r *WebACL_CustomResponseBody) AWSCloudFormationType() string {
	return "AWS::WAFv2::WebACL.CustomResponseBody"
}
