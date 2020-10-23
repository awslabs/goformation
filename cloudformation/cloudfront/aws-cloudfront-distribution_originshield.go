package cloudfront

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// Distribution_OriginShield AWS CloudFormation Resource (AWS::CloudFront::Distribution.OriginShield)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-originshield.html
type Distribution_OriginShield struct {

	// Enabled AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-originshield.html#cfn-cloudfront-distribution-originshield-enabled
	Enabled bool `json:"Enabled"`

	// OriginShieldRegion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-originshield.html#cfn-cloudfront-distribution-originshield-originshieldregion
	OriginShieldRegion string `json:"OriginShieldRegion,omitempty"`

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
func (r *Distribution_OriginShield) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.OriginShield"
}
