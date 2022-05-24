// Code generated by "go generate". Please don't change this file directly.

package rekognition

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// StreamProcessor_S3Destination AWS CloudFormation Resource (AWS::Rekognition::StreamProcessor.S3Destination)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-s3destination.html
type StreamProcessor_S3Destination struct {

	// BucketName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-s3destination.html#cfn-rekognition-streamprocessor-s3destination-bucketname
	BucketName string `json:"BucketName"`

	// ObjectKeyPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-s3destination.html#cfn-rekognition-streamprocessor-s3destination-objectkeyprefix
	ObjectKeyPrefix *string `json:"ObjectKeyPrefix,omitempty"`

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
func (r *StreamProcessor_S3Destination) AWSCloudFormationType() string {
	return "AWS::Rekognition::StreamProcessor.S3Destination"
}
