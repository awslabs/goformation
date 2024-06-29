// Code generated by "go generate". Please don't change this file directly.

package entityresolution

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// IdMappingWorkflow_IntermediateSourceConfiguration AWS CloudFormation Resource (AWS::EntityResolution::IdMappingWorkflow.IntermediateSourceConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-intermediatesourceconfiguration.html
type IdMappingWorkflow_IntermediateSourceConfiguration struct {

	// IntermediateS3Path AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-intermediatesourceconfiguration.html#cfn-entityresolution-idmappingworkflow-intermediatesourceconfiguration-intermediates3path
	IntermediateS3Path string `json:"IntermediateS3Path"`

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
func (r *IdMappingWorkflow_IntermediateSourceConfiguration) AWSCloudFormationType() string {
	return "AWS::EntityResolution::IdMappingWorkflow.IntermediateSourceConfiguration"
}