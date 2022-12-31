// Code generated by "go generate". Please don't change this file directly.

package cloudfront

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// ResponseHeadersPolicy_RemoveHeadersConfig AWS CloudFormation Resource (AWS::CloudFront::ResponseHeadersPolicy.RemoveHeadersConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-removeheadersconfig.html
type ResponseHeadersPolicy_RemoveHeadersConfig struct {

	// Items AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-removeheadersconfig.html#cfn-cloudfront-responseheaderspolicy-removeheadersconfig-items
	Items []ResponseHeadersPolicy_RemoveHeader `json:"Items"`

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
func (r *ResponseHeadersPolicy_RemoveHeadersConfig) AWSCloudFormationType() string {
	return "AWS::CloudFront::ResponseHeadersPolicy.RemoveHeadersConfig"
}
