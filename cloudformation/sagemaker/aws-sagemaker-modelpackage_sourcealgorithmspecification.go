// Code generated by "go generate". Please don't change this file directly.

package sagemaker

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// ModelPackage_SourceAlgorithmSpecification AWS CloudFormation Resource (AWS::SageMaker::ModelPackage.SourceAlgorithmSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-sourcealgorithmspecification.html
type ModelPackage_SourceAlgorithmSpecification struct {

	// SourceAlgorithms AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-sourcealgorithmspecification.html#cfn-sagemaker-modelpackage-sourcealgorithmspecification-sourcealgorithms
	SourceAlgorithms []ModelPackage_SourceAlgorithm `json:"SourceAlgorithms"`

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
func (r *ModelPackage_SourceAlgorithmSpecification) AWSCloudFormationType() string {
	return "AWS::SageMaker::ModelPackage.SourceAlgorithmSpecification"
}
