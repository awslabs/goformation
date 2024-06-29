// Code generated by "go generate". Please don't change this file directly.

package qbusiness

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Retriever_RetrieverConfiguration AWS CloudFormation Resource (AWS::QBusiness::Retriever.RetrieverConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-retrieverconfiguration.html
type Retriever_RetrieverConfiguration struct {

	// KendraIndexConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-retrieverconfiguration.html#cfn-qbusiness-retriever-retrieverconfiguration-kendraindexconfiguration
	KendraIndexConfiguration *Retriever_KendraIndexConfiguration `json:"KendraIndexConfiguration,omitempty"`

	// NativeIndexConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-retrieverconfiguration.html#cfn-qbusiness-retriever-retrieverconfiguration-nativeindexconfiguration
	NativeIndexConfiguration *Retriever_NativeIndexConfiguration `json:"NativeIndexConfiguration,omitempty"`

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
func (r *Retriever_RetrieverConfiguration) AWSCloudFormationType() string {
	return "AWS::QBusiness::Retriever.RetrieverConfiguration"
}
