// Code generated by "go generate". Please don't change this file directly.

package nimblestudio

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// StreamingImage_StreamingImageEncryptionConfiguration AWS CloudFormation Resource (AWS::NimbleStudio::StreamingImage.StreamingImageEncryptionConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-streamingimage-streamingimageencryptionconfiguration.html
type StreamingImage_StreamingImageEncryptionConfiguration struct {

	// KeyArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-streamingimage-streamingimageencryptionconfiguration.html#cfn-nimblestudio-streamingimage-streamingimageencryptionconfiguration-keyarn
	KeyArn *string `json:"KeyArn,omitempty"`

	// KeyType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-streamingimage-streamingimageencryptionconfiguration.html#cfn-nimblestudio-streamingimage-streamingimageencryptionconfiguration-keytype
	KeyType string `json:"KeyType"`

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
func (r *StreamingImage_StreamingImageEncryptionConfiguration) AWSCloudFormationType() string {
	return "AWS::NimbleStudio::StreamingImage.StreamingImageEncryptionConfiguration"
}
