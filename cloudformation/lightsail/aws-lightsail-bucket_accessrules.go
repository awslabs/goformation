package lightsail

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Bucket_AccessRules AWS CloudFormation Resource (AWS::Lightsail::Bucket.AccessRules)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-bucket-accessrules.html
type Bucket_AccessRules struct {

	// AllowPublicOverrides AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-bucket-accessrules.html#cfn-lightsail-bucket-accessrules-allowpublicoverrides
	AllowPublicOverrides *bool `json:"AllowPublicOverrides,omitempty"`

	// GetObject AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-bucket-accessrules.html#cfn-lightsail-bucket-accessrules-getobject
	GetObject *string `json:"GetObject,omitempty"`

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
func (r *Bucket_AccessRules) AWSCloudFormationType() string {
	return "AWS::Lightsail::Bucket.AccessRules"
}
